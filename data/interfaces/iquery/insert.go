package iquery

type Insert interface {
	Query

	To(tableName string) Insert
	OnConflict(query string) Insert
}
