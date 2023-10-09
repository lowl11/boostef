package ef

import (
	"github.com/lowl11/boostef/data/interfaces/imigrate"
)

func Migrate(entities ...imigrate.Entity) {
	for _, entity := range entities {
		found, err := entity.CheckDestination()
		if err != nil {
			panic(err)
		}

		if !found {
			// create table
			if err = entity.CreateDestination(); err != nil {
				panic(err)
			}
		} else {
			// compare columns & decide, call alter table or not
			if err = entity.Compare(); err != nil {
				panic(err)
			}
		}
	}
}
