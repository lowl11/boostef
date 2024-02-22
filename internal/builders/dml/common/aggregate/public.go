package aggregate

import (
	"github.com/lowl11/boostef/internal/helpers/stringc"
)

func (aggregate *Aggregate) String(_ ...string) string {
	return aggregate.condition
}

func (aggregate *Aggregate) Count(field, sign string, value any) {
	aggregate.condition = "COUNT(" + field + ") " + sign + " " + stringc.ToString(value)
}

func (aggregate *Aggregate) Avg(field, sign string, value any) {
	aggregate.condition = "AVG(" + field + ") " + sign + " " + stringc.ToString(value)
}

func (aggregate *Aggregate) Max(field, sign string, value any) {
	aggregate.condition = "MAX(" + field + ") " + sign + " " + stringc.ToString(value)
}

func (aggregate *Aggregate) Min(field, sign string, value any) {
	aggregate.condition = "MIN(" + field + ") " + sign + " " + stringc.ToString(value)
}
