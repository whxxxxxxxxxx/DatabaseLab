package dto

type TokenData struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
