package utils

import ( 
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(getJWTSecret())

func getJWTSecret() string{
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "rahasia_super_aman_2026"
	}

	return secret
}

type Claims struct {
	UserID int `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateTokenI(userID int, email string) (string, error){
	expirationTime := time.Now().Add(24 * time.Hour) 

	claims := &Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssudAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "simple-auth",
			Subject: "user-token",
		},
	}

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error){
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
		return jwtSecret, nil 
	}) 

	if err != nil{
		return nil, err
	}

	if !token.Valid{
		return nil, errors.New("invalid token")
	}
	
	return claims, nil
}

