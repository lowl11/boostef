package migrator

import "github.com/lowl11/boostef/data/interfaces/iquery"

type foundPair struct {
	found bool
	col   iquery.Column
}
