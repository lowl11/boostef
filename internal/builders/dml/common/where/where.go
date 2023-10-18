package where

type Where struct {
	alias      string
	conditions []string
	or         bool
}

func New(alias ...string) *Where {
	where := &Where{
		conditions: make([]string, 0, 5),
	}

	if len(alias) > 0 && len(alias[0]) > 0 {
		where.alias = alias[0]
	}

	return where
}

func NewOr(alias ...string) *Where {
	where := &Where{
		conditions: make([]string, 0, 5),
		or:         true,
	}

	if len(alias) > 0 && len(alias[0]) > 0 {
		where.alias = alias[0]
	}

	return where
}
