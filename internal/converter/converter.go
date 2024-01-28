package converter

import (
	"github.com/lowl11/flex"
	"reflect"
)

type Converter struct {
	entity any
	name   string
}

func New(entity any) *Converter {
	if !flex.Type(reflect.TypeOf(entity)).IsStruct() {
		panic("Given entity is not struct")
	}

	return &Converter{
		entity: entity,
		name:   reflect.TypeOf(entity).Name(),
	}
}
