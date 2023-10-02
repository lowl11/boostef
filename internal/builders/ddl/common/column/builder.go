package column

import "github.com/lowl11/boostef/data/interfaces/iquery"

type Builder struct {
	name     string
	dataType iquery.DataType
}

func New(name ...string) *Builder {
	builder := &Builder{}

	if len(name) > 0 {
		builder.name = name[0]
	}

	return builder
}
