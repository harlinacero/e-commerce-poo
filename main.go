package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ecommerce.com/ecommerce/domain"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var shoppingCar domain.ShoppingCar

	fmt.Println("Bienvenido al ecommerce Majestic")
	user, err := addUser()
	if err != nil {
		fmt.Println("Error al crear el usuario", err)
		return
	}

	fmt.Printf("%s Desea agregar un producto? (s/n)", user.GetUsername())
	addProduct, _ := reader.ReadString('\n')
	addProduct = strings.TrimSpace(addProduct)

	for addProduct == "s" {
		product := addNewProduct()

		fmt.Println("Desea agregar el producto al carrito? (s/n)")

		addProductToShoppingCar, _ := reader.ReadString('\n')
		addProductToShoppingCar = strings.TrimSpace(addProductToShoppingCar)

		if addProductToShoppingCar == "s" {
			shoppingCar.AddProduct(*product)
		}

		fmt.Printf("%s Desea agregar otro producto? (s/n)", user.GetUsername())
		addProduct, _ = reader.ReadString('\n')
		addProduct = strings.TrimSpace(addProduct)
	}

	showShoppingCar(shoppingCar)
}

func showShoppingCar(shoppingCar domain.ShoppingCar) {
	fmt.Println("******************************************************************")
	fmt.Println("Usted ha agregado los siguientes productos a su carrito de compras")
	for i, product := range shoppingCar.Productcs {
		fmt.Println("************ Producto ", i+1, " ************")
		fmt.Println("Nombre: ", product.GetName())
		fmt.Println("Descripción: ", product.GetDescription())
		fmt.Println("Valor Bruto: ", product.GetGrossValue())
		fmt.Println("Impuesto: ", product.GetTaxAddValue())
		fmt.Println("Descuento: ", product.GetDiscValue())
		fmt.Println("Valor Real: ", product.GetRealValue())
	}

	fmt.Println("******************************************************************")
	fmt.Println("Total de la Compra: ", shoppingCar.CalculateTotalValue())
}

func addUser() (*domain.User, error) {
	var user *domain.User
	var name string
	var email string
	var password string
	var password2 string
	var address string

	fmt.Println("Para iniciar, ingresa tus datos")
	fmt.Println("Nombre: ")
	fmt.Scanln(&name)
	fmt.Println("Email: ")
	fmt.Scanln(&email)
	fmt.Println("Contraseña: ")
	fmt.Scanln(&password)
	fmt.Println("Confirma tu contraseña: ")
	fmt.Scanln(&password2)

	if password != password2 {
		return user, errors.New("Contraseñas no coinciden")
	}

	fmt.Println("Dirección: ")
	fmt.Scanln(&address)

	user = domain.NewUser(name, name, email, password, address)
	username := user.GetUsername()
	fmt.Printf("El usuario %s ha sido creado", username)
	fmt.Println("")
	return user, nil
}

func addNewProduct() *domain.Product {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("******************************************************************")
	fmt.Println("Introduce el nombre del producto")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Introduce una descripcion")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Println("Introduce el valor bruto del producto")
	grossvalueStr, _ := reader.ReadString('\n')
	grossvalueStr = strings.TrimSpace(grossvalueStr)
	grossvalue, _ := strconv.ParseFloat(grossvalueStr, 64)

	fmt.Println("Introduce el porcentaje de impuesto")
	taxpercentajeStr, _ := reader.ReadString('\n')
	taxpercentajeStr = strings.TrimSpace(taxpercentajeStr)
	taxpercentaje, _ := strconv.ParseFloat(taxpercentajeStr, 64)

	fmt.Println("Introduce el porcentaje de descuento")
	discpercentajeStr, _ := reader.ReadString('\n')
	discpercentajeStr = strings.TrimSpace(discpercentajeStr)
	discpercentaje, _ := strconv.ParseFloat(discpercentajeStr, 64)

	fmt.Println("Introduce la cantidad de stock")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)
	stock, _ := strconv.Atoi(stockStr)

	product := domain.NewProduct(name, name, description, grossvalue, discpercentaje, taxpercentaje, stock)

	fmt.Println(product.GetName())

	grossvalue = product.GetGrossValue()
	taxvalue := product.GetTaxAddValue()
	discvalue := product.GetDiscValue()
	realvalue := product.GetRealValue()

	fmt.Println("Valor Bruto: ", grossvalue)
	fmt.Println("Impuestos: ", taxvalue)
	fmt.Println("Decuentos: ", discvalue)
	fmt.Println("Total: ", realvalue)

	return product
}
