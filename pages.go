package globals

import "time"

type Page struct {
	ID             int64
	Path           string
	Arguments      string
	Comment        string
	SeoTitle       string
	SeoKeywords    string
	SeoDescription string
	Content        string
	Template       string
	TemplateFile   string
	ContentType    string
	PermissionCode string
	HttpCode       int
	System         bool
	UpdatedAt      time.Time
	CreatedAt      time.Time
}
