package iquery

type Query interface {
	Get(params ...string) string
}
