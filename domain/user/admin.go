package user

import (
	"fmt"

	"github.com/google/uuid"
)

// UserBase is the base struct for the user
type admin struct {
	UserBase // Herencia de usuario
}

// NewAdmin returns new Admin
func NewAdmin(username string, email string, password string) *admin {
	return &admin{
		UserBase{
			id: uuid.New().String(),
			username: username,
			email:    email,
			password: password,
		},
	}
}

// SancitonUser sanciona a un usuario
func (a *admin) SancitonUser(user UserBase, message string) {
	fmt.Println("El usuario ", user.GetUsername(), " ha sido sancionado")
	fmt.Println("Motivo: ", message)
}

// DeleteUser elimina un usuario 
func (a *admin) UnsuscribeUser(username string) {
	//logica para suspender o dar de baja un usuario
	fmt.Println("El usuario ", username, " ha sido dado de baja")
}
