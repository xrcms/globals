package globals

type Cacher interface {
	Get(key string) string
	Set(key, value string)
	Check(key string) bool
	Clear(key string)
}
