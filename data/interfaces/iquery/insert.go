package iquery

type Insert interface {
	Query

	GetParamStatus() (string, bool)
	To(tableName string) Insert
	OnConflict(query string) Insert
}
