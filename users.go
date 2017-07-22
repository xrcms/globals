package globals

import (
	"time"
)

type User struct {
	ID             int64
	Username       string
	Password       string
	Groups         []int64
	Email          string
	EmailConfirmed bool
	FirstName      string
	LastName       string
	SecondName     string
	FullName       string
	Phone          string
	Skype          string
	Site           string
	Sex            string
	Photo          int64
	BirthDay       time.Time
	LastLogin      time.Time
	SecretKey      string
	SessionID      string
	VkID           string
	OkID           string
	MailRuID       string
	YandexID       string
	GoogleID       string
	FacebookID     string
	GithubID       string
	CityID         int64
	About          string
	UpdatedAt      time.Time
	CreatedAt      time.Time
}

type Userser interface {
	Get(userID int64) User
	GetBySocID(socKey, socID string) User
	GetByUsername(username string) User
	GetByEmail(email string) User
	Save(user User) (int64, error)
	GetList(page, limit int64, where, orderBy string) (users []User, count int64)
	Clear(userID int64)
	PrepareUser(row map[string]string) User
	Delete(userID int64) error
	PasswordHash(password string) string
}