package globals

import (
	"time"
)

type Directory struct {
	ID              int64
	Table           string
	Name            string
	UserFields      []Field
	Sort            int
	UpdatedAt       time.Time
	CreatedAt       time.Time
	Elements        Elementser
}

type Directoryer interface {
	GetList() []Directory
	Get(directoryID int64) Directory
	GetByTableName(tableName string) Directory
	Save(directory Directory) (int64, error)
	Delete(directoryID int64) error
	Clear(directoryID int64)
	ClearTables()
}
