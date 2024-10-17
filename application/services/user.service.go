package services

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"ecommerce.com/ecommerce/domain/user"
)

// Agregar nuevo usuario
func AddUser() (*user.Customer, error) {
	reader := bufio.NewReader(os.Stdin)
	var customer *user.Customer
	var name string
	var email string
	var password string
	var password2 string
	var address string

	fmt.Println("Crear usuario vendedor o cliente? (v/c)")
	typeUser, _ := reader.ReadString('\n')
	typeUser = strings.TrimSpace(typeUser)

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
		return customer, errors.New("Contraseñas no coinciden")
	}

	fmt.Println("Dirección: ")
	fmt.Scanln(&address)

	customer = user.NewCustomer(name, email, password, address, "123456")
	username := customer.GetUsername()
	fmt.Printf("El cliente %s ha sido creado", username)

	return customer, nil
}
