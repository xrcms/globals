package globals

type Modules interface {
	Run(module string, params map[string]string) (err error, result []byte)
}
