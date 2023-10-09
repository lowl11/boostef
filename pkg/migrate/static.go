package migrate

import (
	"github.com/lowl11/boostef/data/interfaces/imigrate"
	"github.com/lowl11/boostef/internal/migrator"
)

func Entity(tableName string) imigrate.Entity {
	return migrator.NewEntity(tableName)
}
