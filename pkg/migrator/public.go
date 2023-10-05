package migrator

import (
	"context"
	"github.com/lowl11/boostef/internal/services/entity_service"
	"log"
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
		columns := entity_service.New(ent).Columns()

		givenColumns := make([]column, 0, len(columns))
		for _, col := range columns {
			var notNull bool
			var primary bool
			var foreign bool
			var length int

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
						log.Fatal(err)
					}

					length = l
				} else if strings.Contains(tag, "foreign(") {
					//tag = strings.ReplaceAll(tag, "foreign(", "")
					//tag = strings.ReplaceAll(tag, ")", "")
					foreign = true
				}
			}

			givenColumns = append(givenColumns, column{
				Name:      col.Name,
				Type:      col.DataType,
				Length:    length,
				IsPrimary: primary,
				IsForeign: foreign,
				NotNull:   notNull,
			})
		}

		givenEntities = append(givenEntities, entity{
			columns: givenColumns,
		})
	}

	migrator.givenEntities = givenEntities
	return migrator
}

func (migrator *Migrator) GetReal() *Migrator {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	rows, err := migrator.connection.QueryxContext(ctx, "")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	realEntities := make([]entity, 0, 20)
	//for rows.Next() {
	//	if err = rows.StructScan(); err != nil {
	//		panic(err)
	//	}
	//}

	migrator.realEntities = realEntities
	return migrator
}
