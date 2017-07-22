package globals

import (
	"time"
)

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
