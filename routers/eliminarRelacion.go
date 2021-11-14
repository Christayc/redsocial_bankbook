package routers

import (
	"net/http"

	"github.com/Christayc/redsocial_bankbook/bd"
	"github.com/Christayc/redsocial_bankbook/models"
)

/*EliminarRelacion elimina la relacion entre dos usuarios*/
func EliminarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioRelacionID = ID
	t.UsuarioID = IDUsuario
	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al borrar la relacion", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se pudo borrar la relacion del usuario", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
