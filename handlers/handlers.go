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
	router.HandleFunc("/registro", midlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", midlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", midlew.ChequeoBD(midlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificar-perfil", midlew.ChequeoBD(midlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", midlew.ChequeoBD(midlew.ValidarJWT(routers.InsertarTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", midlew.ChequeoBD(midlew.ValidarJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
