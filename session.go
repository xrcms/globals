package globals

type Sessioner interface {
	SocAuth() bool
	Auth(username, password string, remember bool) bool
	AuthAs(userID int64, remember bool) bool
	Logout()
	User() User
	IsAuth() bool
	IsMobile() bool
}
