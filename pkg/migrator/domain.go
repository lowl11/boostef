package migrator

import "github.com/lowl11/boostef/data/interfaces/iquery"

type entity struct {
	tableName string
	columns   []column
	exist     bool
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

type realEntityColumn struct {
	Table           string  `db:"table_name"`
	Name            string  `db:"column_name"`
	OrdinalPosition int     `db:"ordinal_position"`
	Default         *string `db:"column_default"`
	Nullable        string  `db:"is_nullable"`
	DataType        string  `db:"data_type"`
	MaxLength       *int    `db:"character_maximum_length"`
}

type realEntityMy struct {
	Name     string  `db:"Field"`
	Type     string  `db:"Type"`
	Nullable bool    `db:"Null"`
	Key      *string `db:"Key"`
	Default  *string `db:"Default"`
}

type tableKey struct {
	ConstraintName string `db:"constraint_name"`
	Table          string `db:"table_name"`
	Column         string `db:"column_name"`
}
