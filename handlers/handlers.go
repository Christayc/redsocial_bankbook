package handlers

import (
	"log"
	"net/http"
	"os"

	midlew "github.com/Christayc/redsocial_bankbook/middlew"
	"github.com/Christayc/redsocial_bankbook/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el handler y pongo a escucar el server*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", midlew.ChequeoBD(routers.Registros)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
