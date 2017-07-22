package globals

import "time"

type Elementser interface {
	GetByID(elementID int64) (element Element)
	GetByCode(code string) (element Element)
	GetList(page, limit int64, where, orderBy string, onlyActive bool) (elements []Element, count int64)
	PrepareElement(row map[string]string) Element
	Save(element Element, user User) (int64, error)
	Delete(elementID int64) error
}

type Categorieser interface {
	GetByID(categoryID int64) Category
	GetByCode(code string) Category
	GetList() []Category
	Clear()
	PrepareCategory(row map[string]string) Category
	Save(category Category) (int64, error)
	Delete(categoryID int64) error
}

type Requester interface {
	Host() string
	UserAgent() string
	CheckCaptcha() bool
	Header(key string) string
	UseSSL() bool
	Uri() string
	IP() string
	Keys() []string
	Map() map[string][]string
	File(key string) MultipartFile
	FileArr(key string) []MultipartFile
	Float(key string) float64
	FloatArr(key string) []float64
	Int(key string) int64
	IntArr(key string) []int64
	Bool(key string) bool
	BoolArr(key string) []bool
	String(key string) string
	StringArr(key string) []string
	Time(key string) time.Time
	TimeArr(key string) []time.Time
	SetVal(key string, values []string)
}

type Filter interface {
	Where(request Requester) string
}

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
	UpdatedAt       time.Time
	CreatedAt       time.Time
	Elements        Elementser
	Categories      Categorieser
	Filter          Filter
	IsGoods         bool
	IsOffers        bool
}
