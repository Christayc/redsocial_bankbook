package bd

import (
	"context"
	"time"

	"github.com/Christayc/redsocial_bankbook/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario recibre un email y confirma en la bd si ya existe el usuario*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial_bankbook")
	col := db.Collection("usuario")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {

		return resultado, false, ID
	}
	return resultado, true, ID

}
