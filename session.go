package globals

type Sessioner interface {
	SocAuth() bool
	Auth(username, password string, remember bool) bool
	User() User
	IsAuth() bool
	IsMobile() bool
}
