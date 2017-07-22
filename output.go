package globals

type Outputer interface {
	SetPageFromModule(httpCode int, seoTitle, seoKeywords, seoDescription string)
	GetHeaders() []string
	AddHeader(header string)
	SetPage(page Page)
	GetPage() Page
	SetParams(params map[string]string)
	GetParams() map[string]string
	ModuleParam(paramName, defaultValue string) string
	SetVar(key string, value interface{})
	GetVar(key string) interface{}
	GetAllVars() map[string]interface{}
	AddMessage(message, level string)
	GetMessages() map[string][]string
	AddJS(pathOrCode string)
	GetScripts() []string
	AddCSS(pathOrCode string)
	GetStyles() []string
}
