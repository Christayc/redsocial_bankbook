package midlew

import (
	"net/http"

	"github.com/Christayc/redsocial_bankbook/bd"
)

/*ChequeoBD es el middlew que permite ver el estado de la bd*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {

			http.Error(w, "Conexion perdida con la Base de Datos", 5000)
			return
		}
		next.ServeHTTP(w, r)
	}
}
