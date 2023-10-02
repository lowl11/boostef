package truncate_table

type Builder struct {
	tableName string
}

func New(tableName ...string) *Builder {
	builder := &Builder{}

	if len(tableName) > 0 {
		builder.tableName = tableName[0]
	}

	return builder
}
