package vars

import "github.com/lowl11/boostef/pkg/enums/sqls"

func TableName(sql string) string {
	switch sql {
	case sqls.MySQL:
		return "TABLE_NAME"
	default:
		return "table_name"
	}
}

func Tables(sql string) string {
	switch sql {
	case sqls.MySQL:
		return "information_schema.TABLES"
	default:
		return "information_schema.tables"
	}
}

func Columns(sql string) []string {
	switch sql {
	case sqls.MySQL:
		return []string{}
	default:
		return []string{
			"table_name", "column_name", "data_type", "column_default",
			"is_nullable", "character_maximum_length",
		}
	}
}

func ColumnTable(sql string) string {
	switch sql {
	case sqls.MySQL:
		return ""
	default:
		return "information_schema.columns"
	}
}
