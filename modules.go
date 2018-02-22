package globals

type Modules interface {
	Run(module string, params map[string]string) (result []byte, err error)
}
