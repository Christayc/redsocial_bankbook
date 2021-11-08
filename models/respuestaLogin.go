package models

/* RespuestaLogin tiene el token que devuelve con el loging*/
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
