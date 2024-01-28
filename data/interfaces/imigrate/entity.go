package imigrate

import "github.com/lowl11/boostef/data/interfaces/iquery"

type Entity interface {
	Table() string
	Name() string
	Columns(columns ...iquery.Column) Entity
	PartitionColumns(columns ...string) Entity
	CheckDestination() (bool, error)
	CreateDestination() error
	Compare() error
}
