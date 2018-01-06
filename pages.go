package globals

import (
	"time"
)

type Page struct {
	ID             int64
	Path           string
	Arguments      string
	Comment        string
	SeoTitle       string
	SeoKeywords    string
	SeoDescription string
	Content        []byte
	Template       string
	TemplateFile   string
	ContentType    string
	AddHttpHeaders map[string]string
	PermissionCode string
	HttpCode       int
	System         bool
	GZIP           bool
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type Pageer interface {
	GetByID(pageID int64) Page
	GetList(page, limit int64, where, orderBy string) (pages []Page, count int64)
	PreparePage(row map[string]string) Page
	Save(page Page) (int64, error)
	Delete(pageID int64) error
	GetByURL(url string) (page Page, arguments map[string][]string, err error)
	Clear()
}
