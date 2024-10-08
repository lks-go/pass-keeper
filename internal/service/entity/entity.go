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

type DataLoginPass struct {
	ID       int32
	Title    string
	Login    string
	Password string
}

type DataText struct {
	ID    int32
	Title string
	Text  string
}

type DataCard struct {
	ID      int32
	Title   string
	Number  string
	Owner   string
	ExpDate string
	CVCCode string
}

type DataBinary struct {
	ID    int32
	Title string
	Body  <-chan byte
}
