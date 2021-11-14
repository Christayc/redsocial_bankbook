package bd

import (
	"context"
	"time"

	"github.com/Christayc/redsocial_bankbook/models"
)

/*InsertarRelacion relaciona dos usuarios*/
func InsertarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("redsocial_bankbook")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
