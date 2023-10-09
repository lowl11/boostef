package imigrate

import "github.com/lowl11/boostef/data/interfaces/iquery"

type Entity interface {
	Columns(columns ...iquery.Column) Entity
	CheckDestination() (bool, error)
	CreateDestination() error
	Compare() error
}
