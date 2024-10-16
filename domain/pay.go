package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Pay struct {
	id     string // ID
	method string // Metodo de pago puede ser tarjeta, efectivo, etc
	value  float64 // Valor del pago
	order  Order // Pedido
}

// Constructor
func NewPay(method string, value float64, order Order) *Pay {
	return &Pay{
		id:     uuid.New().String(),
		method: method,
		value:  value,
		order:  order,
	}
}

// Procesamiento del pago
func (p *Pay) ProcessPay() {
	fmt.Println("Procesando pago")

	// Aquí se realizaría la lógica de pago
	time.AfterFunc(2*time.Second, func() {
		p.order.UpdateState("Pagado")
		fmt.Println("Comprobando pago, por favor espere...")
	})

	time.Sleep(3 * time.Second)
	fmt.Println("Pago realizado")
}