package globals

import (
	"time"
)

type Infoblock struct {
	ID              int64
	Table           string
	Code            string
	Name            string
	SeoTitle        string
	SeoKeywords     string
	SeoDescription  string
	UserFields      []Field
	ListPageURL     string
	CategoryPageURL string
	DetailPageURL   string
	Sort            int
	UpdatedAt       time.Time
	CreatedAt       time.Time
	Elements        Elementser
	Categories      Categorieser
	Filter          Filter
	IsGoods         bool
	IsOffers        bool
}

type Infoblocker interface {
	GetList() []Infoblock
	Get(infoblockID int64) Infoblock
	GetByTableName(tableName string) Infoblock
	Save(infoblock Infoblock) (int64, error)
	Delete(infoblockID int64) error
	Clear(infoblockID int64)
	ClearTables()
}

type Filter interface {
	Where(request Requester) string
}
