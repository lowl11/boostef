package drop_index

import "github.com/lowl11/boostef/pkg/enums/sqls"

type Builder struct {
	sql   string
	name  string
	table string
}

func New(name ...string) *Builder {
	builder := &Builder{
		sql: sqls.Postgres,
	}

	if len(name) > 0 {
		builder.name = name[0]
	}

	return builder
}
