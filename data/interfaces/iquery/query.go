package iquery

type Query interface {
	String(params ...string) string
}
