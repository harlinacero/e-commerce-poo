package order

import (
	"time"

	"ecommerce.com/ecommerce/domain/product"
	"ecommerce.com/ecommerce/domain/user"
	"github.com/google/uuid"
)

type Order struct {
	id       string // ID
	user     *user.Customer // Usuario
	products []product.Product // Lista de Productos
	date     time.Time // Fecha
	state    string // Estado del pedido
	total    float64 // Valor total del pedido
}

// Constructor
func NewOrder(user *user.Customer, products []product.Product, state string, total float64) *Order {
	return &Order{
		id:       uuid.New().String(),
		user:     user,
		products: products,
		state:    state,
		total:    total,
		date:     time.Now(),
	}
}

// Actualiza el estado del pedido
func (o *Order) UpdateState(state string) {
	o.state = state
}

// Obtiene el estado del pedido
func (o *Order) GetState() string {
	return o.state
}

// Cancela el pedido
func (o *Order) CancelOrder() {
	o.state = "Cancelado"
}

// Obtiene el pedido
func (o *Order) GetOrder() *Order {
	return o
}

// Obtiene el total del pedido
func (o *Order) GetTotal() float64 {
	return o.total
}