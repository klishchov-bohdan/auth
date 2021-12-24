package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const TokenLifeTime = 10
const Secret = "access_secret"
const RefreshSecret = "refresh_secret"
const RefreshTokenLifeTime = 60

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/profile", Profile)
	http.HandleFunc("/refresh", Refresh)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		req := new(LoginRequest)
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := NewUserRepo().GetByEmail(req.Email)
		if err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}
		// authenticated

		tokenString, err := GenerateToken(user.ID, TokenLifeTime, Secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		refreshString, err := GenerateToken(user.ID, RefreshTokenLifeTime, RefreshSecret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := &LoginResponse{
			AccessToken:  tokenString,
			RefreshToken: refreshString,
		}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
	}

}

func Profile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tokenString := GetTokenFromBearerString(r.Header.Get("Authorization"))
		claims, err := ValidateToken(tokenString, Secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		user, err := NewUserRepo().GetByID(claims.ID)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		resp := UserResponseProfile{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
	}

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)

	default:
		http.Error(w, "Only GET method", http.StatusMethodNotAllowed)
	}

}
