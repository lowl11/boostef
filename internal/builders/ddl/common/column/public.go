package column

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"strings"
)

func (builder *Builder) Get(params ...string) string {
	if len(params) == 0 {
		return ""
	}

	query := strings.Builder{}
	query.WriteString(builder.name)
	query.WriteString(" ")
	builder.dataType.Write(params[0], &query)
	return strings.TrimSpace(strings.ReplaceAll(query.String(), "% FIELD_NAME %", builder.name))
}

func (builder *Builder) Name(name string) iquery.Column {
	builder.name = name
	return builder
}

func (builder *Builder) GetName() string {
	return builder.name
}

func (builder *Builder) DataType(dataType iquery.DataType) iquery.Column {
	builder.dataType = dataType
	return builder
}

func (builder *Builder) GetDataType() iquery.DataType {
	return builder.dataType
}
