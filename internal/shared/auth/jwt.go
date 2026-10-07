package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func getSecretPass() []byte {
	secret := os.Getenv("PassJWT")

	if secret == "" {
		log.Fatal("secret harus di isi")
	}
	return []byte(secret)
}

type JwtKlaim struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJwt(id int, username, role string) (string, error) {
	godotenv.Load()

	now := time.Now().UTC()
	waktu := now.Add(10 * time.Hour)
	claimjwt := &JwtKlaim{
		Id:       id,
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tamplate_golang",
			ExpiresAt: jwt.NewNumericDate(waktu),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimjwt)

	pass, err := token.SignedString(getSecretPass())
	if err != nil {
		return "", err
	}

	return pass, nil

}

func ValidateJwt(tokenString string) (*JwtKlaim, error) {
	claim := &JwtKlaim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (any, error) {

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {

			return nil, fmt.Errorf("salah metode %v", t.Header["alg"])

		}

		return getSecretPass(), nil

	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claim, nil
}
