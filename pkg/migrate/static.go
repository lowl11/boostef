package migrate

import (
	"github.com/lowl11/boostef/data/interfaces/imigrate"
	"github.com/lowl11/boostef/internal/converter"
	"github.com/lowl11/boostef/internal/migrator"
)

func Entity(schema, tableName string) imigrate.Entity {
	return migrator.NewEntity(schema, tableName)
}

func Convert(entity any) imigrate.Entity {
	return converter.
		New(entity).
		Entity()
}

func Run(scripts ...string) {
	//
}
