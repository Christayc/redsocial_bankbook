package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncriptarPassword es la rutina que permite encryptar constaseñas*/
func EncriptarPassword(pass string) (string, error) {
	costo := 10
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
