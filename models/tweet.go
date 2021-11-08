package models

/*Tweet estructura de un tweet*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
