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
		http.Error(w, "Datos incorrectos", http.StatusBadRequest)
		return
	}

	var status bool

	status, err = bd.ModificarRegisitro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error al actualizar el usuario", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se actualizo los datos", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
