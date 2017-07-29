package globals

import (
	"time"
)

type Element struct {
	ID              int64
	InfoblockID     int64
	Code            string
	Name            string
	PreviewPicture  int64
	PreviewText     string
	DetailPicture   int64
	DetailText      string
	ActiveFrom      time.Time
	ActiveFromNull  bool
	ActiveTo        time.Time
	ActiveToNull    bool
	Active          bool
	AuthorID        int64
	EditorID        int64
	Categories      []int64
	SeoTitle        string
	SeoKeywords     string
	SeoDescription  string
	UserFields      []Field
	ListPageURL     string
	CategoryPageURL string
	DetailPageURL   string
	UpdatedAt       time.Time
	CreatedAt       time.Time

	// Значения для инфоблоков товаров
	Offers     []int64
	OffersList []Element
	Quantity   int64
	Price      float64
	PriceMin   float64
	PriceMax   float64
}

type Elementser interface {
	GetByID(elementID int64) (element Element)
	GetByCode(code string) (element Element)
	GetList(page, limit int64, where, orderBy string, onlyActive bool) (elements []Element, count int64)
	PrepareElement(row map[string]string) Element
	Save(element Element, user User) (int64, error)
	Delete(elementID int64) error
}
