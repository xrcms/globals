package globals

import "time"

type Permission struct {
	ID          int64
	Code        string
	Prefix      string
	Description string
	Groups      []int64
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type Permissionser interface {
	Check(keys string, groups []int64) (bool, error)
	List() []Permission
	RemoveGroup(key string, groupID int64) error
	AddGroup(key string, groupID int64) error
	Save(permission Permission) (int64, error)
	PreparePermission(row map[string]string) Permission
}
