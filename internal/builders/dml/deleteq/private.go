package deleteq

import "github.com/lowl11/boostef/data/interfaces/iquery"

func (builder *Builder) applyWhere(whereFunc func(builder iquery.Where)) *Builder {
	whereFunc(builder.where)
	return builder
}
