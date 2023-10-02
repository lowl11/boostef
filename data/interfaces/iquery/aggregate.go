package iquery

type Aggregate interface {
	Count(field, sign string, value any)
	Avg(field, sign string, value any)
	Max(field, sign string, value any)
	Min(field, sign string, value any)
}
