package globals

import (
	"time"
)

type Category struct {
	ID int64
	InfoblockID int64
	Code string
	Name string
	Level int
	ParentID int64
	SeoTitle string
	SeoKeywords string
	SeoDescription string
	ListPageURL     string
	CategoryPageURL string
	Sort int
	UpdatedAt time.Time
	CreatedAt time.Time
}
