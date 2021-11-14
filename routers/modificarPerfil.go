package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Christayc/redsocial_bankbook/bd"
	"github.com/Christayc/redsocial_bankbook/models"
)

/*ModificarPerfil modifica el perfil*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al actualizar el usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se actualizo los datos", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
