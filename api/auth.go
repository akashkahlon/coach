package api

import (
	"coach/services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type Credentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func LoginHandler(db *gorm.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var creds Credentials
			if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}
	
			if creds.Email == "" || creds.Password == "" {
				http.Error(w, "Email and password are required", http.StatusBadRequest)
				return
			}

			log.Println("creds: ", creds)
			log.Println("db: ", db)
			token, tokenExpirationTime, err := services.LoginUser(db, creds.Email, creds.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token,
				Expires:  tokenExpirationTime,
				HttpOnly: true,
			})
	
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Login successful"))
		}
}