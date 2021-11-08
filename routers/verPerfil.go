package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Christayc/redsocial_bankbook/bd"
)

/*VerPerfil visualizar un perfil de usuario*/
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID es obligatorio", http.StatusBadRequest)
	}
	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error en la consulta", http.StatusBadRequest)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(perfil)
}
