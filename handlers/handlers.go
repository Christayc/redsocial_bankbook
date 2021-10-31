package handlers

import(
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

)
/*Manejadores seteo mi puerto, el handler y pongo a escucar el server*/
func Manejadores(
	router := mux.NetRouter()

	PORT := os.Getenv("PORT")
	if PORT== ""  {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
)