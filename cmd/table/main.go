package main

import (
	"fmt"
	"github.com/lowl11/boostef/pkg/builder"
	"github.com/lowl11/boostef/pkg/ctypes"
	"github.com/lowl11/boostef/pkg/enums/sqls"
)

func main() {
	create()
	//drop()
}

func create() {
	fmt.Println("=== QUERY ===")
	fmt.Println(builder.
		CreateTable("products").
		Sql(sqls.Postgres).
		IfNotExist().
		Column(
			// boolean
			builder.Column("some_bool").DataType(ctypes.Bool()),
			builder.Column("some_boolean").DataType(ctypes.Boolean()),

			// text
			builder.Column("tiny_text").DataType(ctypes.TinyText()),
			builder.Column("text").DataType(ctypes.Text()),
			builder.Column("varchar").DataType(ctypes.Varchar(123)),
			builder.Column("char").DataType(ctypes.Char(123)),
			builder.Column("character").DataType(ctypes.Character(123)),
			builder.Column("charvar").DataType(ctypes.CharacterVarying(123)),
			builder.Column("id").DataType(ctypes.Serial().Primary().NotNull().AutoIncrement()),

			// numbers
			builder.Column("small_int").DataType(ctypes.SmallInt()),
			builder.Column("middle_integer").DataType(ctypes.Integer()),
			builder.Column("big_int").DataType(ctypes.Bigint()),

			builder.Column("small_serial").DataType(ctypes.SmallSerial()),
			builder.Column("serial").DataType(ctypes.Serial()),
			builder.Column("big_serial").DataType(ctypes.BigSerial()),

			builder.Column("real").DataType(ctypes.Real()),

			// time
			builder.Column("timestamp").DataType(ctypes.Timestamp()),

			builder.Column("brand_id").DataType(ctypes.Uuid().Foreign("brand.id")),
		).
		Get())
}

func drop() {
	//fmt.Println()
}
