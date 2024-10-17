package services

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ecommerce.com/ecommerce/domain/product"
)

// Agregar nuevo producto
func AddNewProduct() product.Product {
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

	fmt.Printf("El producto %s es digital o físico? (d/f)", name)
	typeProduct, _ := reader.ReadString('\n')
	typeProduct = strings.TrimSpace(typeProduct)

	var product product.Product

	switch typeProduct {
	case "d":
		product = addDigitalProduct(name, description, grossvalue, discpercentaje, taxpercentaje, stock)
	case "f":
		product = addPhysicalProduct(name, description, grossvalue, discpercentaje, taxpercentaje, stock)
	default:
		fmt.Println("Tipo de producto no valido")
	}

	return product
}

func addPhysicalProduct(name, description string, grossvalue, discpercentaje, taxpercentaje float64, stock int) product.Product {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Introduce el peso del producto")
	weightStr, _ := reader.ReadString('\n')
	weightStr = strings.TrimSpace(weightStr)
	weight, _ := strconv.ParseFloat(weightStr, 64)

	fmt.Println("Introduce la altura del producto")
	heightStr, _ := reader.ReadString('\n')
	heightStr = strings.TrimSpace(heightStr)
	height, _ := strconv.ParseFloat(heightStr, 64)

	fmt.Println("Introduce el ancho del producto")
	widthStr, _ := reader.ReadString('\n')
	widthStr = strings.TrimSpace(widthStr)
	width, _ := strconv.ParseFloat(widthStr, 64)

	fmt.Println("Introduce la longitud del producto")
	lengthStr, _ := reader.ReadString('\n')
	lengthStr = strings.TrimSpace(lengthStr)
	length, _ := strconv.ParseFloat(lengthStr, 64)

	product := product.NewPhysicalProduct(name, name, description, grossvalue, discpercentaje, taxpercentaje, stock, weight, height, width, length)
	
	return product
}

func addDigitalProduct(name, description string, grossvalue, discpercentaje, taxpercentaje float64, stock int) product.Product {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Introduce el formato del archivo")
	fileFormat, _ := reader.ReadString('\n')
	fileFormat = strings.TrimSpace(fileFormat)

	fmt.Println("Introduce el tamaño del archivo")
	sizeStr, _ := reader.ReadString('\n')
	sizeStr = strings.TrimSpace(sizeStr)
	size, _ := strconv.ParseFloat(sizeStr, 64)

	product := product.NewDigitalProduct(name, name, description, grossvalue, discpercentaje, taxpercentaje, stock, fileFormat, size)
	return product
}
