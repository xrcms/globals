package globals

import "time"

type Comment struct {
	ID              int64
	Text            string
	Active          bool
	AuthorID        int64
	AuthorName      string
	EditorID        int64
	Code            string
	ReplyToDate     time.Time
	UpdatedAt       time.Time
	CreatedAt       time.Time
}
