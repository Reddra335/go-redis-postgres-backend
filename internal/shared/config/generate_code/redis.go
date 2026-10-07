package generatecode

import (
	"crypto/rand"
	"encoding/hex"
	"newProjek/internal/shared/logger"

	"github.com/google/uuid"
)

func UnixID() string {
	log := logger.New()
	u := make([]byte, 32)

	_, err := rand.Read(u)
	if err != nil {
		log.WithError(err).Fatal("Pembuatan id gagal")
	}
	res := hex.EncodeToString(u)

	return res
}

func UUID() string {
	log := logger.New()
	u, err := uuid.NewV7()
	if err != nil {
		log.WithError(err).Fatal("gagal membuat id")
	}
	return u.String()

}
