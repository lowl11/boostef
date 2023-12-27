package ef

import (
	"context"
	"github.com/lowl11/boost/log"
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

func Run(ctx context.Context, scripts ...string) error {
	for _, script := range scripts {
		if err := Execute(ctx, script); err != nil {
			return err
		}
	}

	return nil
}

func RunError(ctx context.Context, onError func(err error), scripts ...string) {
	for _, script := range scripts {
		if err := Execute(ctx, script); err != nil {
			onError(err)
		}
	}
}

func MustRun(ctx context.Context, scripts ...string) {
	if err := Run(ctx, scripts...); err != nil {
		log.Fatal(err, "Run scripts error")
	}
}
