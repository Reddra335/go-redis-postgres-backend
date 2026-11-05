package generateenkripsi

import (
	"crypto/rand"
	"encoding/hex"
	"newProjek/internal/shared/logger"

	"github.com/sirupsen/logrus"
)

func IdRds() string {
	log := logger.New()
	ge := make([]byte, 30)

	_, err := rand.Read(ge)

	if err != nil {
		log.WithFields(logrus.Fields{

			"operation": "generate enkripsi",
			"layer":     "generate enkripsi",
		}).WithError(err).Fatal("gagal membuat keamanan")
		return ""
	}

	data := hex.EncodeToString(ge)
	return data
}
