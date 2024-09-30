package model

type Token struct {
	Token     string `json:"token"`
	ExpiredIn int    `json:"expire_in"` // seconds
}

type AuthResponse struct {
	AccessToken  Token `json:"access_token"`
	RefreshToken Token `json:"-"`
}
