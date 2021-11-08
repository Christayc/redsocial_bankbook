package bd

import (
	"github.com/Christayc/redsocial_bankbook/models"
	"golang.org/x/crypto/bcrypt"
)

/*intentoLogin realiza la validacion con la base de datos.*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {

		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {

		return usu, false
	}
	return usu, true
}
