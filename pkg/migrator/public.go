package migrator

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/helpers/stringc"
	"github.com/lowl11/boostef/internal/services/entity_service"
	"github.com/lowl11/boostef/internal/services/executor"
	"github.com/lowl11/boostef/internal/system"
	"github.com/lowl11/boostef/pkg/builder"
	"strconv"
	"strings"
	"time"
)

func (migrator *Migrator) SetEntities(entities ...any) *Migrator {
	if len(entities) == 0 {
		return migrator
	}

	/*
		1. Take info from given entities & save it
		2. Take info from real database & save it
		3. Compare given & real entities
		4. Update different "things" from given to real database
	*/

	/*
		entity info:
			- column names
			- column types
			- column lengths
			- column is not null
			- column default values

			- index...
	*/

	// take info from given entities & save it
	givenEntities := make([]entity, 0, len(entities))
	for _, ent := range entities {
		es := entity_service.New(ent)
		columns := es.Columns()
		tableName := es.TableName()

		givenColumns := make([]column, 0, len(columns))
		for _, col := range columns {
			var notNull bool
			var primary bool
			var foreign bool
			var length int
			var defValue string

			for _, tag := range col.EfTags {
				switch tag {
				case "nn":
					notNull = true
				case "primary":
					primary = true
				}

				if strings.Contains(tag, "len(") {
					tag = strings.ReplaceAll(tag, "len(", "")
					tag = strings.ReplaceAll(tag, ")", "")
					l, err := strconv.Atoi(tag)
					if err != nil {
						panic(err)
					}

					length = l
				} else if strings.Contains(tag, "foreign(") {
					//tag = strings.ReplaceAll(tag, "foreign(", "")
					//tag = strings.ReplaceAll(tag, ")", "")
					foreign = true
				} else if strings.Contains(tag, "default") {
					_, after, _ := strings.Cut(tag, ":")
					defValue = after
				}
			}

			dt := col.DataType
			if notNull {
				dt.NotNull()
			}

			if primary {
				dt.Primary()
			}

			if len(defValue) > 0 {
				dt.Default(defValue)
			}

			givenColumns = append(givenColumns, column{
				Name:      col.Name,
				Type:      dt,
				Length:    length,
				IsPrimary: primary,
				IsForeign: foreign,
				NotNull:   notNull,
			})
		}

		givenEntities = append(givenEntities, entity{
			tableName: tableName,
			columns:   givenColumns,
		})
	}

	migrator.givenEntities = givenEntities
	return migrator
}

func (migrator *Migrator) GetReal() *Migrator {
	realEntities := make([]entity, 0, 20)
	for _, given := range migrator.givenEntities {
		tableName := given.tableName

		// check, if table exists in tables list
		exist, err := executor.Exist(
			migrator.connection,
			system.GetTable(migrator.dialect, given.tableName),
		)
		if err != nil {
			panic(err)
		}

		// if it does not exist, add but with flag exist: false
		if !exist {
			realEntities = append(realEntities, entity{
				tableName: tableName,
			})
			continue
		}

		// get columns for this (iteration) table from database
		realColumns, err := executor.GetStruct[realEntityColumn](
			migrator.connection,
			system.GetColumns(migrator.dialect, tableName),
		)
		if err != nil {
			panic(err)
		}

		// collect columns
		columns := make([]column, 0, 10)
		for _, col := range realColumns {
			var length int
			if col.MaxLength != nil {
				length = *col.MaxLength
			}

			var isPrimary bool
			for _, key := range migrator.keys {
				if key.Table == tableName && key.Column == col.Name {
					isPrimary = true
				}
			}

			columns = append(columns, column{
				Name:         col.Name,
				Length:       length,
				IsPrimary:    isPrimary,
				NotNull:      col.Nullable == "NO",
				DefaultValue: stringc.ToString(col.Default),
			})
		}

		// add entity info from database
		realEntities = append(realEntities, entity{
			tableName: tableName,
			columns:   columns,
			exist:     true,
		})
	}

	migrator.realEntities = realEntities
	return migrator
}

func (migrator *Migrator) Migrate() {
	for index, realEntity := range migrator.realEntities {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		defer cancel()

		// if entity does not exist
		if !realEntity.exist {
			// create entity
			columns := make([]iquery.Column, 0, len(realEntity.columns))

			// migrator.givenEntities[index] - this is shit code
			// because we are inside real entities iteration, but using its index
			// for given entities list
			for _, col := range migrator.givenEntities[index].columns {
				columns = append(columns, builder.Column(col.Name).DataType(col.Type))
			}

			q := builder.
				CreateTable(realEntity.tableName).
				Sql(migrator.dialect).
				IfNotExist().
				Column(columns...).
				Get()

			_, err := migrator.connection.ExecContext(ctx, q)
			if err != nil {
				panic(err)
			}

			continue
		}

		// compare columns
		for _, realCol := range realEntity.columns {
			for _, givenCol := range migrator.givenEntities[index].columns {
				if realCol.Name == givenCol.Name {
					//
				}
			}
		}

	}
}
