package services

import (
	"coach/models"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey []byte
var tokenExpirationTime time.Time

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func loadConfig() {
	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	expiryHours, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRY_HOURS"))
	if err != nil {
		log.Fatalf("Invalid JWT_TOKEN_EXPIRY_HOURS: %v", err)
	}
	tokenExpirationTime = time.Now().Add(time.Duration(expiryHours) * time.Hour)
}

func createToken(email string) (string, error) {
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return tokenString, nil
}

func LoginUser(db *gorm.DB, email string, password string) (string, time.Time, error) {
	loadConfig()
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", tokenExpirationTime, errors.New("invalid email or password")
		}
		return "", tokenExpirationTime, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", tokenExpirationTime, errors.New("invalid email or password")
	}

	tokenString, err := createToken(email)
	if err != nil {
		return "", tokenExpirationTime, err
	}

	return tokenString, tokenExpirationTime, nil
}