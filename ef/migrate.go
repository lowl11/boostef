package ef

import (
	"context"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boostef/data/interfaces/imigrate"
)

func Migrate(entities ...imigrate.Entity) error {
	for _, entity := range entities {
		found, err := entity.CheckDestination()
		if err != nil {
			return err
		}

		if !found {
			// create table
			if err = entity.CreateDestination(); err != nil {
				return err
			}
		} else {
			// compare columns & decide, call alter table or not
			if err = entity.Compare(); err != nil {
				return err
			}
		}
	}

	return nil
}

func MustMigrate(entities ...imigrate.Entity) {
	if err := Migrate(entities...); err != nil {
		panic(err)
	}
}

func MigrateError(handler func(string, error), entities ...imigrate.Entity) {
	for _, entity := range entities {
		found, err := entity.CheckDestination()
		if err != nil {
			if handler != nil {
				handler(entity.Name(), err)
			}
			continue
		}

		if !found {
			// create table
			if err = entity.CreateDestination(); err != nil {
				if handler != nil {
					handler(entity.Name(), err)
				}
				continue
			}
		} else {
			// compare columns & decide, call alter table or not
			if err = entity.Compare(); err != nil {
				if handler != nil {
					handler(entity.Name(), err)
				}
				continue
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
