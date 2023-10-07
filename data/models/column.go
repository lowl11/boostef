package models

import "github.com/lowl11/boostef/data/interfaces/iquery"

type Column struct {
	Name     string
	DataType iquery.DataType
	EfTags   []string
	TableTag string
}
