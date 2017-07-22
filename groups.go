package globals

import "time"

type Group struct {
	ID          int64
	Name        string
	Description string
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type Groupser interface {
	Get(groupID int64) Group
	GetList(page, limit int64, where, orderBy string) (groups []Group, count int64)
	PrepareGroup(row map[string]string) Group
	Save(group Group) (int64, error)
	Delete(groupID int64) error
	Clear(groupID int64)
}
