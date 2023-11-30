package create_index

type Builder struct {
	name       string
	table      string
	unique     bool
	columns    []string
	ifNotExist bool
}

func New(name ...string) *Builder {
	builder := &Builder{
		columns: []string{},
	}

	if len(name) > 0 {
		builder.name = name[0]
	}

	return builder
}
