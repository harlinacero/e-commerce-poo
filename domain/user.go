package domain

import "github.com/google/uuid"

type User struct {
	id       string
	username string
	email    string
	password string
	address  string
}

func NewUser(username string, email string,
	password string, address string) *User {
	return &User{
		id:       uuid.New().String(),
		username: username,
		email:    email,
		password: password,
		address:  address,
	}
}

func (u *User) SetUsername(username string) {
	u.username = username
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetAddress(address string) {
	u.address = address
}

func (u *User) GetUsername() string {
	return u.username
}
func (u *User) GetEmail() string {
	return u.email
}
func (u *User) GetPassword() string {
	return u.password
}
func (u *User) GetAddress() string {
	return u.address
}
