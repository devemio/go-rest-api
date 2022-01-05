package users

type User struct {
	Id           int
	Username     string
	EmailAddress string
}

func New() *User {
	return &User{}
}
