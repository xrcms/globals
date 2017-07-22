package globals

import "time"

type EmailTemplate struct {
	ID          int64
	Name        string
	Subject     string
	Body        string
	System      bool
	UpdatedAt   time.Time
	CreatedAt   time.Time
}

type EmailTemplateser interface {
	GetList(page, limit int64, orderBy string) (templates []EmailTemplate, count int64)
	Get(templateID int64) EmailTemplate
	PrepareTemplate(row map[string]string) (template EmailTemplate)
	Save(template EmailTemplate) (int64, error)
	Delete(templateID int64) error
}
