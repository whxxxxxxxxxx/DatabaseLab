package dto

type TokenData struct {
	Token        string      `json:"token"`
	RefreshToken string      `json:"refresh_token"`
	User         interface{} `json:"user"`
}
