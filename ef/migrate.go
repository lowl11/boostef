package ef

import (
	"context"
	"github.com/lowl11/boostef/data/interfaces/imigrate"
	"log"
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

func Run(ctx context.Context, scripts ...string) {
	for _, script := range scripts {
		if err := Execute(ctx, script); err != nil {
			log.Fatal(err)
		}
	}
}
