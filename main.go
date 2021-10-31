package main

import (
	"log"

	"github.com/Christayc/redsocial_bankbook/bd"
	"github.com/Christayc/redsocial_bankbook/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexion con la Base de Datos")
		return

	}
	handlers.Manejadores()

}
