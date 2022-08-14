package users

type User struct {
	ID           int64
	Username     string
	EmailAddress string
	Images       []Image
}

type Image struct {
	ID  int64
	Url string
}
