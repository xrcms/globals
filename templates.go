package globals

type Templateser interface {
	Load(tplName string, data interface{}) (result string)
	ParseTemplate(tpl, tplName string, data interface{}) (result string)
}
