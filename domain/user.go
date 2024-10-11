package domain

type User struct {
	ID       string
	Username string
	Email    string
	Password string
	Address  string
}

func NewUser(id string, username string, email string,
	password string, address string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
		Address:  address,
	}
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) SetAddress(address string) {
	u.Address = address
}

func (u *User) GetUsername() string {
	return u.Username
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetPassword() string {
	return u.Password
}
func (u *User) GetAddress() string {
	return u.Address
}
