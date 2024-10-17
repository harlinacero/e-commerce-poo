package user

import (
	"ecommerce.com/ecommerce/domain/product"
	"github.com/google/uuid"
)

// UserBase is the base struct for the user
type Seller struct {
	UserBase // Herencia de usuario
	phone    string
}

// Constructor
func NewSeller(username string, email string, password string, address string, phone string) *Seller {
	return &Seller{
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
func (s *Seller) SetPhone(phone string) {
	s.phone = phone
}

// Getters
func (s *Seller) GetPhone() string {
	return s.phone
}

func (s *Seller) UpdateProductDisccount(product *product.Product, disccount float64) {
	(*product).SetDisccountPercentage(disccount)
}

func (s *Seller) UpdateProductPrice(product *product.Product, price float64) {
	(*product).SetGrossValue(price)
}

func (s *Seller) UpdateProductStock(product *product.Product, stock int) {
	(*product).SetStock(stock)
}