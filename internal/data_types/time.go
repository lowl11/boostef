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
	Interval (Postgres)
*/

func Timestamp() iquery.DataType {
	return createCustom(all.Timestamp)
}

func TimestampZ() iquery.DataType {
	return createCustom(postgres.TimestampZ).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.TimestampZ
		}

		return all.Timestamp
	})
}

func SmallDateTime() iquery.DataType {
	return createCustom(mssql.SmallDateTime).setCustom(func(sql string) string {
		switch sql {
		case sqls.MSSQL:
			return mssql.SmallDateTime
		}

		return mysql.DateTime
	})
}

func DateTime() iquery.DataType {
	return createCustom(mysql.DateTime).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return all.Timestamp
		}

		return mysql.DateTime
	})
}

func DateTimeOffset() iquery.DataType {
	return createCustom(mssql.DateTimeOffset).setCustom(func(sql string) string {
		switch sql {
		case sqls.MSSQL:
			return mssql.DateTimeOffset
		}

		return all.Timestamp
	})
}

func DateTime2() iquery.DataType {
	return createCustom(mssql.DateTime2).setCustom(func(sql string) string {
		switch sql {
		case sqls.MSSQL:
			return mssql.DateTime2
		}

		return all.Timestamp
	})
}

func Date() iquery.DataType {
	return createCustom(all.Date)
}

func Time() iquery.DataType {
	return createCustom(all.Time)
}

func Year() iquery.DataType {
	return createCustom(mysql.Year).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.Year
		}

		return all.Integer
	})
}
