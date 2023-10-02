package where

type Where struct {
	conditions []string
	or         bool
}

func New() *Where {
	return &Where{
		conditions: make([]string, 0, 5),
	}
}

func NewOr() *Where {
	return &Where{
		conditions: make([]string, 0, 5),
		or:         true,
	}
}
