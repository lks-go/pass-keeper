package entity

const (
	AuthTokenHeader     = "auth_token"
	UserLoginHeaderName = "User-Login"
)

type User struct {
	ID           string
	Login        string
	Password     string
	PasswordHash string
}
