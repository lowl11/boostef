package data_types

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types/all"
	"github.com/lowl11/boostef/internal/data_types/mssql"
	"github.com/lowl11/boostef/internal/data_types/mysql"
	"github.com/lowl11/boostef/internal/data_types/postgres"
	"github.com/lowl11/boostef/pkg/enums/sqls"
)

/*
	Numeric (Postgres, MS SQL)
	Double Precision (Postgres)
*/

func Bit(size ...int) iquery.DataType {
	return createCustom(mysql.Bit, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return all.SmallInt
		}

		return mysql.Bit
	})
}

func TinyInt(size ...int) iquery.DataType {
	return createCustom(mysql.TinyInt, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return all.SmallInt
		}

		return mysql.TinyInt
	})
}

func SmallInt(size ...int) iquery.DataType {
	return createCustom(all.SmallInt, size...)
}

func Int(size ...int) iquery.DataType {
	return createCustom(mssql.Int, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return all.Integer
		}

		return mssql.Int
	})
}

func MediumInt(size ...int) iquery.DataType {
	return createCustom(mysql.MediumInt, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return all.Integer
		case sqls.MSSQL:
			return mssql.Int
		}

		return mysql.MediumText
	})
}

func Integer(size ...int) iquery.DataType {
	return createCustom(all.Integer, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MSSQL:
			return mssql.Int
		}

		return all.Integer
	})
}

func Bigint(size ...int) iquery.DataType {
	return createCustom(all.Bigint, size...)
}

func Real(size ...int) iquery.DataType {
	return createCustom(all.Real, size...)
}

func Decimal(size ...int) iquery.DataType {
	return createCustom(all.Decimal, size...)
}

func Dec(size ...int) iquery.DataType {
	return createCustom(all.Decimal, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.Dec
		}

		return all.Decimal
	})
}

func Double(size ...int) iquery.DataType {
	return createCustom(mysql.Double, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.Double
		}

		return all.Real
	})
}

func Float(size ...int) iquery.DataType {
	return createCustom(mysql.Float, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return all.Real
		}

		return mysql.Float
	})
}

func SmallMoney(size ...int) iquery.DataType {
	return createCustom(mssql.SmallMoney, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.Money
		case sqls.MySQL:
			return all.Real
		}

		return mssql.SmallMoney
	})
}

func Money(size ...int) iquery.DataType {
	return createCustom(postgres.Money, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return all.Real
		}

		return postgres.Money
	})
}

func SmallSerial() iquery.DataType {
	return createCustom(postgres.SmallSerial).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.SmallSerial
		}

		return all.TinyInt
	})
}

func Serial() iquery.DataType {
	return createCustom(postgres.Serial).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.Serial
		}

		return all.Integer + " AUTO_INCREMENT"
	})
}

func BigSerial() iquery.DataType {
	return createCustom(postgres.BigSerial).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.BigSerial
		}

		return all.Bigint
	})
}
