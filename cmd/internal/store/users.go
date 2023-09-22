package store

type User struct {
	Username string `binding: "required,min=5,max=30"`
	Password string `binding: "required,min=4,max=30"`
}

var Users []*User
