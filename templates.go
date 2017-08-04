package globals

type Templateser interface {
	AddTemplate(tplName, tpl string)
	Load(tplName string, data interface{}) (result string)
	ParseTemplate(tpl, tplName string, data interface{}) (result string)
}

