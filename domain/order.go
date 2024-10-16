package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	id       string // ID
	user     *User // Usuario
	products []Product // Lista de Productos
	date     time.Time // Fecha
	state    string // Estado del pedido
	total    float64 // Valor total del pedido
}

// Constructor
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