package migrator

import (
	"context"
	"fmt"
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
	table   string
	columns []iquery.Column
}

func NewEntity(table string) *Entity {
	return &Entity{
		table:   table,
		columns: ef.EntityColumns(),
	}
}

func (entity *Entity) Columns(columns ...iquery.Column) imigrate.Entity {
	entity.columns = columns
	return entity
}

func (entity *Entity) CheckDestination() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	result, err := ef.ExecuteResult(ctx, builder.
		Select().
		From(entity.table).
		Limit(1).
		Get())
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			return false, nil
		}
		return false, err
	}

	return len(result) > 0, nil
}

func (entity *Entity) CreateDestination() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := ef.Execute(ctx, builder.
		CreateTable(entity.table).IfNotExist().
		Column(entity.columns...).
		Sql(ef_core.Get().SQL()).
		Get())
	if err != nil {
		return err
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
			where.Equal("table_name", entity.table)
		}).
		Get())
	if err != nil {
		return err
	}

	destinationColumns := make([]iquery.Column, 0, len(entity.columns))
	for _, result := range results {
		if len(result) == 0 {
			continue
		}

		col := builder.Column(fmt.Sprintf("%s", result["column_name"]))

		// data_type
		var defaultValue *string
		if result["column_default"] != nil {
			v := fmt.Sprintf("%s", result["column_default"])
			defaultValue = &v
		}

		var maxLength int64
		if result["character_maximum_length"] != nil {
			maxLength = result["character_maximum_length"].(int64)
		}

		col.DataType(convertDataType(
			fmt.Sprintf("%s", result["data_type"]),
			fmt.Sprintf("%s", result["is_nullable"]),
			defaultValue,
			maxLength,
		))

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
						Set("NOT NULL").
						Get())
				case compares.Type:
					newQueries = append(newQueries, builder.
						AlterTable(entity.table).
						AlterColumn(sourceColumn.GetName()).
						Type(sourceColumn.GetDataType()).
						Get())
				}
			}

			//fmt.Println(sourceColumn.GetName(), different)
		}
	}

	for key, value := range founds {
		if value.found {
			continue
		}

		newQueries = append(newQueries, builder.
			AlterTable(entity.table).
			AddColumn(key).
			Type(value.col.GetDataType()).
			Get())
	}

	// execute queries
	for _, newQuery := range newQueries {
		if err = ef.Execute(ctx, newQuery); err != nil {
			panic(err)
		}
	}

	return nil
}

func convertDataType(dataType, isNullable string, defaultValue *string, maxLength int64) iquery.DataType {
	var dt iquery.DataType

	switch dataType {
	case "text":
		dt = ctypes.Text()
	case "varchar", "character varying":
		if maxLength == 0 {
			return nil
		}

		dt = ctypes.Varchar(int(maxLength))
	case "uuid":
		dt = ctypes.Uuid()
	case "timestamp", "timestampz", "timestamp without time zone", "timestamp with time zone":
		dt = ctypes.Timestamp()
	case "int", "integer":
		dt = ctypes.Integer()
	case "serial":
		dt = ctypes.Serial()
	case "boolean", "bool":
		dt = ctypes.Boolean()
	default:
		panic("does not support data type: " + dataType)
	}

	if isNullable == "NO" {
		dt.NotNull()
	}

	if defaultValue != nil {
		dt.Default(*defaultValue)
	}

	return dt
}
