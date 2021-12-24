package main

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponseProfile struct {
	ID    int
	Email string
	Name  string
}
