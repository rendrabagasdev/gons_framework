package models

type Users struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func init() {
	RegisterModel(&Users{})
}
