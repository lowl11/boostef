package postgres

// boolean
const (
	Boolean = "BOOLEAN"
)

// text
const (
	Char             = "CHAR"
	Character        = "CHARACTER"
	CharacterVarying = "CHARACTER VARYING"
	Varchar          = "VARCHAR"
	Text             = "TEXT"
	UUID             = "UUID"
)

// numbers
const (
	SmallInt        = "SMALLINT"
	Integer         = "INTEGER"
	BigInt          = "BIGINT"
	Decimal         = "DECIMAL"
	Numeric         = "NUMERIC"
	DoublePrecision = "DOUBLE PRECISION"
	SmallSerial     = "SMALLSERIAL"
	Serial          = "SERIAL"
	BigSerial       = "BIGSERIAL"
	Money           = "Money"
)

// binary
const (
	Bytea = "BYTEA"
)

// time
const (
	Timestamp  = "TIMESTAMP"
	TimestampZ = "TIMESTAMPZ"
	Date       = "DATE"
	Time       = "TIME"
	Interval   = "INTERVAL"
)
