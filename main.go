package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ecommerce.com/ecommerce/application/services"
	"ecommerce.com/ecommerce/domain/shoppingcar"
)

func main() {

	fmt.Println("Bienvenido al ecommerce Majestic")
	user, err := services.AddUser()
	if err != nil {
		fmt.Println("Error al crear el usuario", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var shoppingCar shoppingcar.ShoppingCar

	fmt.Printf("%s Desea agregar un producto? (s/n)", user.GetUsername())
	addProduct, _ := reader.ReadString('\n')
	addProduct = strings.TrimSpace(addProduct)

	for addProduct == "s" {
		product := services.AddNewProduct()

		fmt.Println("Desea agregar el producto al carrito? (s/n)")

		addProductToShoppingCar, _ := reader.ReadString('\n')
		addProductToShoppingCar = strings.TrimSpace(addProductToShoppingCar)

		if addProductToShoppingCar == "s" {
			shoppingCar.AddProduct(product)
		}

		fmt.Printf("%s Desea crear otro producto? (s/n)", user.GetUsername())
		addProduct, _ = reader.ReadString('\n')
		addProduct = strings.TrimSpace(addProduct)
	}

	shoppingCar.ShowShoppingCar()

}

