package auth

import (
	"errors"
	"fmt"
	"log"
	"newProjek/internal/shared/helper/errprogram"
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
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJwt(ID int, username, email string) (string, error) {
	godotenv.Load()
	now := time.Now().UTC()
	waktu := now.Add(10 * time.Hour)
	claimjwt := &JwtKlaim{
		ID:       ID,
		Username: username,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tamplate_golang",
			ExpiresAt: jwt.NewNumericDate(waktu),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimjwt)

	pass, err := token.SignedString(getSecretPass())
	if err != nil {
		errprogram.ErrorField("error", "GenerateJwt", "Auth", "pembuat JWT gagal", err)
		return "", err
	}

	return pass, nil

}

func ValidateJwt(tokenString string) (*JwtKlaim, error) {
	claim := &JwtKlaim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (any, error) {

		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {

			errprogram.ErrorField("error", "ValidateJWT", "Auth", "Metode Salah", errors.New("Metode salah"))

			return nil, fmt.Errorf("salah metode %v", t.Header["alg"])

		}

		return getSecretPass(), nil

	})

	if err != nil {
		errprogram.ErrorField("error", "ValidateJwt", "Auth", "Error pada pemeriksaan token", err)

		return nil, err
	}
	if !token.Valid {
		errprogram.ErrorField("error", "ValidateJwt", "Auth", "pembuat JWT gagal", errors.New("token tidak valid"))
		return nil, errors.New("token tidak valid")
	}

	return claim, nil
}
