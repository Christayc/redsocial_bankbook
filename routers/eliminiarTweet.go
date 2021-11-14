package routers

import (
	"net/http"

	"github.com/Christayc/redsocial_bankbook/bd"
)

/*EliminarTweet remueve un tweet*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}

	err := bd.BoroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar el tweet", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
