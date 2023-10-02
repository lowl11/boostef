package aggregate

type Aggregate struct {
	condition string
}

func New() *Aggregate {
	return &Aggregate{}
}
