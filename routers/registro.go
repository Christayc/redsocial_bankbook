package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Christayc/redsocial_bankbook/bd"
	"github.com/Christayc/redsocial_bankbook/models"
)

/* Registro es la func para crear en la bd el registro del usuario*/

func Registros(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos ingresados"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {

		http.Error(w, "El email es requerido", 400)
		return
	}
	if len(t.Password) < 6 {

		http.Error(w, "Debes de especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe uan cuenta con este email, por favor colocar una diferente", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
