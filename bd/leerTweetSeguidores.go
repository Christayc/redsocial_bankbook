package bd

import (
	"context"
	"time"

	"github.com/Christayc/redsocial_bankbook/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeerTweetSeguidores recupera mensajes de seguidores*/
func LerTweetSeguidor(ID string, pagina int) ([]models.RespuestaTweetSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("redsocial_bankbook")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	// filtra usuario
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})

	// join para el agregate
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "tweet",
		}})

	// no utiliza maestro
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})

	// ordena descendente con -1 y ascendente con 1
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})

	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)

	var result []models.RespuestaTweetSeguidores
	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true

}
