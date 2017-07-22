package globals

type Optionser interface {
	Get(key string) string
	Save(key, value string) error
	Clear(key string)
}
