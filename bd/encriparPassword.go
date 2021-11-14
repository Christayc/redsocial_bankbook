package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword encripta una contraseña*/
func EncriptarPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(bytes), err
}
