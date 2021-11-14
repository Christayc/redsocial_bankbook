package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Christayc/redsocial_bankbook/bd"
)

/*LeerTweetRelacion recupera mensajes de la relacion*/
func LeerTweetRelacion(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "La pagina es obligatoria", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Pagina es numerico y obligatorio", http.StatusBadRequest)
		return
	}
	respuesta, status := bd.LerTweetSeguidor(IDUsuario, pagina)
	if status == false {
		http.Error(w, "Error al leer los mensajes", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respuesta)
}
