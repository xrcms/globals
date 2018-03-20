package globals

import (
	"time"
)

type DirectoryItem struct {
	ID              int64
	DirectoryID     int64
	Name            string
	UserFields      []Field
	UpdatedAt       time.Time
	CreatedAt       time.Time
}

type DirectoryItemer interface {
	GetByID(itemID int64) (item DirectoryItem)
	GetList(page, limit int64, where, orderBy string) (item []DirectoryItem, count int64)
	PrepareElement(row map[string]string, userFields map[string][]string) DirectoryItem
	Save(item DirectoryItem, user User) (int64, error)
	Delete(itemID int64) error
}
