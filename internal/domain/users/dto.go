package users

type UserInDto struct {
	Username     string `json:"username,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
}
