package main

import (
	"fmt"

	"ecommerce.com/ecommerce/domain"
)

func main() {
	fmt.Println("Bienvenido a la clase de POO")

	addProduct()


}

func addProduct() {
	var name string
	var description string
	var grossvalue float64
	var taxpercentaje float64
	var discpercentaje float64
	var stock int

	fmt.Println("Introduce el nombre del producto")	
	fmt.Scanln(&name)
	fmt.Println("Introduce una descripcion")
	fmt.Scanln(&description)
	fmt.Println("Introduce el valor bruto del producto")
	fmt.Scanln(&grossvalue)
	fmt.Println("Introduce el porcentaje de impuesto")
	fmt.Scanln(&taxpercentaje)
	fmt.Println("Introduce el porcentaje de descuento")
	fmt.Scanln(&discpercentaje)
	fmt.Println("Introduce la cantidad de stock")
	fmt.Scanln(&stock)


	product1 := domain.NewProduct(name, name, description, grossvalue, discpercentaje, taxpercentaje, stock)

	fmt.Println(product1.GetName())

    grossvalue = product1.GetGrossValue()
	taxvalue := product1.GetTaxAddValue()
	discvalue := product1.GetDiscValue()
	realvalue := product1.GetRealValue()
	fmt.Println("Gross Value: ", grossvalue)
	fmt.Println("Tax Value: ", taxvalue)
	fmt.Println("Disc Value: ", discvalue)
	fmt.Println("Real Value: ", realvalue)
}
