package migrator

import (
	"context"
	"fmt"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boostef/data/interfaces/imigrate"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/ef"
	"github.com/lowl11/boostef/internal/compares"
	"github.com/lowl11/boostef/internal/ef_core"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/ctypes"
	"strings"
	"time"
)

type Entity struct {
	schema           string
	table            string
	name             string
	columns          []iquery.Column
	partitionColumns []string
}

func NewEntity(schema, table, name string) *Entity {
	return &Entity{
		schema:  schema,
		table:   table,
		name:    name,
		columns: ef.EntityColumns(),
	}
}

func (entity *Entity) Table() string {
	return entity.table
}

func (entity *Entity) Name() string {
	if entity.name == "" {
		return entity.table
	}

	return entity.name
}

func (entity *Entity) Columns(columns ...iquery.Column) imigrate.Entity {
	entity.columns = columns
	return entity
}

func (entity *Entity) PartitionColumns(columns ...string) imigrate.Entity {
	entity.partitionColumns = columns
	return entity
}

func (entity *Entity) CheckDestination() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := fmt.Sprintf(`
SELECT * FROM information_schema.tables
WHERE table_schema = '%s' AND table_name = '%s'
`, entity.schema, entity.table)

	result, err := ef.ExecuteResult(ctx, query)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			return false, nil
		}

		return false, errors.
			New("Check destination table exist or not").
			SetType("EF_Migrate_CheckDestinationError").
			SetError(err).
			AddContext("table", entity.table)
	}

	return len(result) > 0, nil
}

func (entity *Entity) CreateDestination() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	createTableQuery := builder.
		CreateTable(entity.table).
		IfNotExist().
		Column(entity.columns...).
		Sql(ef_core.Get().SQL())

	if len(entity.partitionColumns) > 0 {
		createTableQuery.PartitionBy(entity.partitionColumns...)
	}

	err := ef.Execute(ctx, createTableQuery.String())
	if err != nil {
		return errors.
			New("Create destination table").
			SetType("EF_Migrate_CreateDestinationTableError").
			SetError(err).
			AddContext("table", entity.table)
	}

	return nil
}

func (entity *Entity) Compare() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	results, err := ef.ExecuteResult(ctx, builder.
		Select("column_name", "data_type", "column_default", "is_nullable", "character_maximum_length").
		From("information_schema.columns").
		Where(func(where iquery.Where) {
			where.
				Equal("table_name", entity.table).
				NotEqual("table_schema", "information_schema")
		}).
		String())
	if err != nil {
		return errors.
			New("String table columns error").
			SetType("EF_Migrate_GetTableColumnsError").
			SetError(err).
			AddContext("table", entity.table)
	}

	indexes, err := ef.ExecuteResult(ctx, builder.
		Select().
		From("pg_indexes").
		Where(func(where iquery.Where) {
			where.Equal("tablename", entity.table)
		}).
		String())
	if err != nil {
		return errors.
			New("String table indices error").
			SetType("EF_Migrate_GetTableIndicesError").
			SetError(err).
			AddContext("table", entity.table)
	}

	destinationColumns := make([]iquery.Column, 0, len(entity.columns))
	for _, result := range results {
		if len(result) == 0 {
			continue
		}

		name := fmt.Sprintf("%s", result["column_name"])
		col := builder.Column(name)
		col.DataType(convertDataType(name, result, indexes))

		destinationColumns = append(destinationColumns, col)
	}

	founds := make(map[string]*foundPair, len(entity.columns))
	for _, col := range entity.columns {
		founds[col.GetName()] = &foundPair{
			found: false,
			col:   col,
		}
	}

	newQueries := make([]string, 0)
	for _, sourceColumn := range entity.columns {
		for _, destinationColumn := range destinationColumns {
			if sourceColumn.GetName() != destinationColumn.GetName() {
				continue
			}

			founds[sourceColumn.GetName()].found = true

			different := sourceColumn.GetDataType().Equals(destinationColumn.GetDataType())
			for _, diff := range different {
				switch diff {
				case compares.NotNullAdd:
					newQueries = append(newQueries, builder.
						AlterTable(entity.table).
						AlterColumn(sourceColumn.GetName()).
						SQL(ef_core.Get().SQL()).
						Set("NOT NULL").
						String())
				case compares.NotNullRemove:
					newQueries = append(newQueries, builder.
						AlterTable(entity.table).
						AlterColumn(sourceColumn.GetName()).
						SQL(ef_core.Get().SQL()).
						Drop("NOT NULL").
						String())
				case compares.Type:
					newQueries = append(newQueries, builder.
						AlterTable(entity.table).
						AlterColumn(sourceColumn.GetName()).
						SQL(ef_core.Get().SQL()).
						Type(sourceColumn.GetDataType()).
						String())
				case compares.UniqueAdd:
					indexName := strings.Builder{}
					_, _ = fmt.Fprintf(&indexName, "%s_%s_unique", entity.table, sourceColumn.GetName())
					newQueries = append(newQueries, builder.
						CreateIndex(indexName.String()).
						Unique().
						TableColumns(entity.table, sourceColumn.GetName()).
						String())
				case compares.UniqueRemove:
					indexName := strings.Builder{}
					_, _ = fmt.Fprintf(&indexName, "%s_%s_unique", entity.table, sourceColumn.GetName())
					newQueries = append(newQueries, builder.
						DropIndex(indexName.String()).
						SQL(ef_core.Get().SQL()).
						Table(entity.table).
						String())
				}
			}
		}
	}

	for key, value := range founds {
		if value.found {
			continue
		}

		newQueries = append(newQueries, builder.
			AlterTable(entity.table).
			AddColumn(key).
			SQL(ef_core.Get().SQL()).
			Type(value.col.GetDataType()).
			String())
	}

	// execute queries
	for _, newQuery := range newQueries {
		if err = ef.Execute(ctx, newQuery); err != nil {
			return errors.
				New("Execute new table error").
				SetType("EF_Migrate_NewTableError").
				SetError(err).
				AddContext("table", entity.table).
				AddContext("query", newQuery)
		}
	}

	return nil
}

func convertDataType(name string, column map[string]any, indexes []map[string]any) iquery.DataType {
	var isUnique bool

	// check indexes
	for _, index := range indexes {
		for key, value := range index {
			valueStr := fmt.Sprintf("%s", value)
			if key == "indexdef" {
				if !strings.Contains(valueStr, "("+name+")") {
					continue
				}

				if strings.Contains(valueStr, "UNIQUE INDEX") {
					isUnique = true
				}
			}
		}
	}

	dataType := fmt.Sprintf("%s", column["data_type"])
	isNullable := fmt.Sprintf("%s", column["is_nullable"])

	// default value
	var defaultValue *string
	if column["column_default"] != nil {
		v := fmt.Sprintf("%s", column["column_default"])
		defaultValue = &v
	}

	// max length
	var maxLength int64
	if column["character_maximum_length"] != nil {
		maxLength = column["character_maximum_length"].(int64)
	}

	var dt iquery.DataType

	switch dataType {
	case "text", "date":
		dt = ctypes.Text()
	case "varchar", "character varying":
		if maxLength == 0 {
			return nil
		}

		dt = ctypes.Varchar(int(maxLength))
	case "uuid":
		dt = ctypes.Uuid()
	case "timestamp", "timestampz",
		"timestamp without time zone", "timestamp with time zone":
		dt = ctypes.Timestamp()
	case "int", "integer", "bigint", "smallint":
		dt = ctypes.Integer()
	case "real":
		dt = ctypes.Real()
	case "serial":
		dt = ctypes.Serial()
	case "boolean", "bool":
		dt = ctypes.Boolean()
	case "jsonb":
		dt = ctypes.JsonB()
	default:
		panic("unsupported data type: " + dataType + ". Name: " + name)
	}

	if isNullable == "NO" {
		dt.NotNull()
	}

	if defaultValue != nil {
		dt.Default(*defaultValue)
	}

	if isUnique {
		dt.Unique()
	}

	return dt
}
