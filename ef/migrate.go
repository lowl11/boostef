package ef

import (
	"github.com/lowl11/boostef/internal/services/ef_core"
	"github.com/lowl11/boostef/pkg/migrator"
)

func Migrate(entities ...any) {
	migrator.
		New(ef_core.Get().Connection(), ef_core.Get().Dialect(), ef_core.Get().Schema()).
		SetEntities(entities...).
		GetReal().
		Migrate()
}
