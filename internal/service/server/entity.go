package server

type User struct {
	ID           string
	Login        string
	Password     string
	PasswordHash string
}
