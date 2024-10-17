package user

// UserBase is the base struct for the user
type UserBase struct {
	id       string // ID
	username string // Username
	email    string // Email
	password string // Password
	address  string 
}


// Setters
func (u *UserBase) SetUsername(username string) {
	u.username = username
}

func (u *UserBase) SetEmail(email string) {
	u.email = email
}

func (u *UserBase) SetPassword(password string) {
	u.password = password
}

func (u *UserBase) SetAddress(address string) {
	u.address = address
}

// Getters
func (u *UserBase) GetUsername() string {
	return u.username
}
func (u *UserBase) GetEmail() string {
	return u.email
}

func (u *UserBase) GetAddress() string {
	return u.address
}
