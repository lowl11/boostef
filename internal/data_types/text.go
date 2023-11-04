package data_types

import (
	"github.com/lowl11/boostef/data/interfaces/iquery"
	"github.com/lowl11/boostef/internal/data_types/all"
	"github.com/lowl11/boostef/internal/data_types/mysql"
	"github.com/lowl11/boostef/internal/data_types/postgres"
	"github.com/lowl11/boostef/pkg/enums/sqls"
)

/*
	Image (MS SQL)

	TinyBlob (MySQL)
	Blob (MySQL)
	MediumBlob (MySQL)
	LongBlob (MySQL)
	Enum (MySQL)
	Set (MySQL)

	Bytea (Postgres)
*/

func Binary(size ...int) iquery.DataType {
	return createCustom(mysql.Binary, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL, sqls.MSSQL:
			return mysql.Binary
		}

		return all.Varchar
	})
}

func VarBinary(size ...int) iquery.DataType {
	return createCustom(mysql.Binary, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL, sqls.MSSQL:
			return mysql.VarBinary
		}

		return all.Varchar
	})
}

func TinyText(size ...int) iquery.DataType {
	return createCustom(mysql.TinyText, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.TinyText
		}

		return all.Text
	})
}

func Text(size ...int) iquery.DataType {
	return createCustom(all.Text, size...)
}

func MediumText(size ...int) iquery.DataType {
	return createCustom(mysql.MediumText, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.MediumText
		}

		return all.Text
	})
}

func LongText(size ...int) iquery.DataType {
	return createCustom(mysql.LongText, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.LongText
		}

		return all.Text
	})
}

func Varchar(size ...int) iquery.DataType {
	return createCustom(all.Varchar, size...)
}

func Char(size ...int) iquery.DataType {
	return createCustom(all.Char, size...)
}

func NChar(size ...int) iquery.DataType {
	return createCustom(mysql.NChar, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.NChar
		}

		return all.Char
	})
}

func NVarchar(size ...int) iquery.DataType {
	return createCustom(mysql.NVarchar, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.NVarchar
		}

		return all.Varchar
	})
}

func NText(size ...int) iquery.DataType {
	return createCustom(mysql.NText, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.MySQL:
			return mysql.NText
		}

		return all.Text
	})
}

func Character(size ...int) iquery.DataType {
	return createCustom(all.Char, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.Character
		}

		return postgres.Varchar
	})
}

func CharacterVarying(size ...int) iquery.DataType {
	return createCustom(all.Varchar, size...).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.CharacterVarying
		}

		return all.Varchar
	})
}

func UUID() iquery.DataType {
	return createCustom(postgres.UUID).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.UUID
		}

		return postgres.Varchar + "(16)"
	})
}

func JsonB() iquery.DataType {
	return createCustom(postgres.JsonB).setCustom(func(sql string) string {
		switch sql {
		case sqls.Postgres:
			return postgres.JsonB
		}

		return all.Text
	})
}
