package join

type Join struct {
	joinType   string
	table      string
	alias      string
	joinColumn string
	mainColumn string
}

func New(joinType string) *Join {
	return &Join{
		joinType: joinType,
	}
}
