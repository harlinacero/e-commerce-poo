package user

import (
	"github.com/google/uuid"
)

type Customer struct {
	UserBase    // Herencia de usuario
	phone       string
}

// Constructor
func NewCustomer(username string, email string, password string, address string, phone string) *Customer {
	return &Customer{
		UserBase: UserBase{
			id:       uuid.New().String(),
			username: username,
			email:    email,
			password: password,
			address:  address,
		},
		phone: phone,
	}
}

// Setters
func (c *Customer) SetPhone(phone string) {
	c.phone = phone
}

// Getters
func (c *Customer) GetPhone() string {
	return c.phone
}
