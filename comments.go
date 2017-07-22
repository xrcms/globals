package globals

import (
	"time"
	"github.com/xrcms/globals"
)

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

type Commentser interface {
	GetList(page, limit int64, where, orderBy string, onlyActive bool) (comments []Comment, count int64)
	Get(commentID int64) Comment
	PrepareComment(row map[string]string) Comment
	Save(comment Comment, user globals.User) (int64, error)
	Delete(commentID int64) error
}