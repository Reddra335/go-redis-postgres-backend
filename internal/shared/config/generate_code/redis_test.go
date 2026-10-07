package generatecode

import (
	"newProjek/internal/shared/logger"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestDisplay(t *testing.T) {
	log := logger.New()
	uji := "k"

	if uji == "" {
		log.WithFields(logrus.Fields{

			"operation": "create otp",
			"layer":     "Redis",
		}).Fatal("kosong")
	}

	log.Info("Berhasil")
}
