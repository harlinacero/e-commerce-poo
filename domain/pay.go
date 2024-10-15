package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Pay struct {
	id     string
	method string
	value  float64
	order  Order
}

func NewPay(method string, value float64, order Order) *Pay {
	return &Pay{
		id:     uuid.New().String(),
		method: method,
		value:  value,
		order:  order,
	}
}

func (p *Pay) ProcessPay() {
	fmt.Println("Procesando pago")

	// Aquí se realizaría la lógica de pago
	time.AfterFunc(2*time.Second, func() {
		p.order.updateState("Pagado")
		fmt.Println("Comprobando pago, por favor espere...")
	})

	time.Sleep(3 * time.Second)
	fmt.Println("Pago realizado")
}