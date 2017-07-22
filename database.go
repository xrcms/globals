package globals

import "database/sql"

type DBField struct {
	Name    string
	Type    string
	NotNull bool
	HasDef  bool
	Num     int
}

type DataBaseResulter interface {
	GetLastInsertID() int64
	FetchRow() map[string]string
	Fetch() bool
	Row() map[string]string
	Free()
}

type DataBaser interface {
	DBName() string
	Close()
	Native() *sql.DB
	Connect() error
	Query(query string) DataBaseResulter
	Get(selectQuery string) []map[string]string
	InsertOnDublicateUpdate(tableName, where string, data map[string]interface{}) (int64, error)
	DeferInsert(tableName string, data map[string]interface{}, max int)
	InsertAllDefered()
	Insert(tableName string, data map[string]interface{}) (int64, error)
	Update(tableName, where string, data map[string]interface{}) (err error)
	Delete(tableName, where string) (err error)
	GetDatabases() []string
	GetEnumType(typeName string) []string
	GetFields(tableName string) []DBField
	GetTables() []string
	DateConvert(s string) string
	ForSQL(str string) string
	Split(s string) []string
	Trim(s string) string
	StripSlashes(s string) string
	AddSlashes(s string) string
	ToTimestamp(value interface{}) string
}
