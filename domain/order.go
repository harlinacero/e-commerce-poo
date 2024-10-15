package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id       string
	user     *User
	products []Product
	date     time.Time
	state    string
	total    float64
}

func NewOrder(user *User, products []Product, state string, total float64) *Order {
	return &Order{
		id:       uuid.New().String(),
		user:     user,
		products: products,
		state:    state,
		total:    total,
		date:     time.Now(),
	}
}

func (o *Order) updateState(state string) {
	o.state = state
}

func (o *Order) GetState() string {
	return o.state
}

func (o *Order) CancelOrder() {
	o.state = "Cancelado"
}

func (o *Order) GetOrder() *Order {
	return o
}

func (o *Order) GetTotal() float64 {
	return o.total
}