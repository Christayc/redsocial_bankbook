package routers

import (
	"errors"
	"strings"

	"github.com/Christayc/redsocial_bankbook/bd"
	"github.com/Christayc/redsocial_bankbook/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email email del usuario*/
var Email string

/*IDUsuario id del usuario*/
var IDUsuario string

/*ProcesarToken procesa el token jwt*/
func ProcesarToken(tk string) (*models.Claim, bool, string, error) {
	miclave := []byte("GoProgra4")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token invalido1")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miclave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido2")
	}
	return claims, false, string(""), err
}
