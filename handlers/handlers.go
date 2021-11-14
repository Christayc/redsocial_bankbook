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
	router.HandleFunc("/verPerfil", midlew.ChequeoBD(midlew.ValidarJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", midlew.ChequeoBD(midlew.ValidarJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/guardar-tweet", midlew.ChequeoBD(midlew.ValidarJWT(routers.InsertarTweet))).Methods("POST")
	router.HandleFunc("/leet-tweet", midlew.ChequeoBD(midlew.ValidarJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminar-tweet", midlew.ChequeoBD(midlew.ValidarJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/alta-relacion", midlew.ChequeoBD(midlew.ValidarJWT(routers.AgregaRelacion))).Methods("POST")
	router.HandleFunc("/baja-relacion", midlew.ChequeoBD(midlew.ValidarJWT(routers.EliminarRelacion))).Methods("DELETE")
	router.HandleFunc("/consulta-relacion", midlew.ChequeoBD(midlew.ValidarJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/lista-usuarios", midlew.ChequeoBD(midlew.ValidarJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/lista-tweet", midlew.ChequeoBD(midlew.ValidarJWT(routers.LeerTweetRelacion))).Methods("GET")
	router.HandleFunc("/subir-avatar", midlew.ChequeoBD(midlew.ValidarJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subir-banner", midlew.ChequeoBD(midlew.ValidarJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtener-avatar", midlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/obtener-banner", midlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
