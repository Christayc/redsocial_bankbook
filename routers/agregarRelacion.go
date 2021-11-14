package routers

import (
	"github.com/Christayc/redsocial_bankbook/bd"
	"github.com/Christayc/redsocial_bankbook/models"
	"net/http"
)

/*AgregarRelacion agrega la relacion a un usuario*/
func AgregaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioRelacionID = ID
	t.UsuarioID = IDUsuario
	status, err := bd.InsertarRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al seguir a usuario", http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se pudo relacionar al usuario", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
