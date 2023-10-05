package migrator

import "github.com/lowl11/boostef/data/interfaces/iquery"

type entity struct {
	columns []column
}

type column struct {
	Name         string
	Type         iquery.DataType
	Length       int
	IsPrimary    bool
	IsForeign    bool
	NotNull      bool
	DefaultValue string
}

type realEntity struct {
	//
}
