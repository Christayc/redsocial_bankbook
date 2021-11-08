package midlew

import (
	"net/http"

	"github.com/Christayc/redsocial_bankbook/routers"
)

/*ValidarJWT verifica jwt*/
func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesarToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token invalido", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
