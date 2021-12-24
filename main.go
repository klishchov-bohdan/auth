package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
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

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only Post method", http.StatusMethodNotAllowed)
	}

}
