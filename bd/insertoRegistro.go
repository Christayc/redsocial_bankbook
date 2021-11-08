package bd

import (
	"context"
	"time"

	"github.com/Christayc/redsocial_bankbook/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro es el paso final que se tiene para ingresar los datos al usuario en la bd*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("redsocial_bankbook")
	col := db.Collection("usuario")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}
