package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RespuestaTweetSeguidores modelo de respuesta de listado de mensajes de seguidores*/
type RespuestaTweetSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId,omitempty"`
	Tweet             struct {
		ID      string    `bson:"_id" json:"_id,omitempty"`
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
	}
}
