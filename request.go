package main

import "github.com/golang-jwt/jwt"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	Token *jwt.Token
}
