package globals

type Cookieser interface {
	Get(name string) string
	Set(name string, value string, maxAge int)
}
