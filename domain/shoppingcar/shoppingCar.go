package shoppingcar

import (
	"fmt"

	"ecommerce.com/ecommerce/domain/product"
	"ecommerce.com/ecommerce/domain/user"
)

type ShoppingCar struct {
	user      user.Customer
	products []product.Product
}

// NewShoppingCar returns new ShoppoingCar.
func NewShoppingCar(user user.Customer, products []product.Product) *ShoppingCar {
	return &ShoppingCar{
		user:      user,
		products: products,
	}
}

// Add product to shopping car
func (sc *ShoppingCar) AddProduct(product product.Product) {
	sc.products = append(sc.products, product)
}

// Remove product of shopping car
func (sc *ShoppingCar) RemoveProduct(productName string) {
	for i, product := range sc.products {
		if (product).GetName() == productName {
			sc.products = append(sc.products[:i], sc.products[i+1:]...)
		}
	}
}

func (sc *ShoppingCar) GetProducts() []product.Product {
	return sc.products
}

func (sc *ShoppingCar) CalculateTotalValue() float64 {
	var total float64
	for _, product := range sc.products {
		total += (product).GetRealValue()
	}

	return total
}

// Muestra el carrito de compras
func (sc *ShoppingCar) ShowShoppingCar() {
	fmt.Println("******************************************************************")
	fmt.Println("Usted ha agregado los siguientes productos a su carrito de compras")
	for i, product := range sc.GetProducts() {
		fmt.Println("************ Producto ", i+1, " ************")
		fmt.Println("Nombre: ", (product).GetName())
		fmt.Println("Descripción: ", (product).GetDescription())

		fmt.Println("Información: ", (product).ShowInfo())

		fmt.Println("Valor Bruto: ", (product).GetGrossValue())
		fmt.Println("Impuesto: ", (product).GetTaxAddValue())
		fmt.Println("Descuento: ", (product).GetDiscValue())
		fmt.Println("Valor Real: ", (product).GetRealValue())
	}

	fmt.Println("******************************************************************")
	fmt.Println("Total de la Compra: ", sc.CalculateTotalValue())
}
