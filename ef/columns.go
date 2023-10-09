package ef

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/ctypes"
)

func EntityColumns() []iquery.Column {
	return []iquery.Column{
		builder.Column("id").DataType(ctypes.Uuid().Primary().NotNull()),
		builder.Column("created_at").DataType(ctypes.Timestamp().NotNull().Default("NOW()")),
		builder.Column("created_at").DataType(ctypes.Timestamp().NotNull().Default("NOW()")),
	}
}
