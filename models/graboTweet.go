package models

import "time"

/*GraboTweets nos permite regalizar el guardado de cada uno de los tweet creados por el usuario*/
type GrabooTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
