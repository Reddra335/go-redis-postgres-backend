package errprogram

import (
	"newProjek/internal/shared/logger"

	"github.com/sirupsen/logrus"
)

func ErrorField(mode string, operation string, layer string, msg string, err error) {

	log := logger.New()
	if "error" == mode {
		log.WithFields(logrus.Fields{
			"operation": operation,
			"layer":     layer,
		}).WithError(err).Error(msg)

	} else if "fatal" == mode {
		log.WithFields(logrus.Fields{
			"operation": operation,
			"layer":     layer,
		}).WithError(err).Error(msg)
	}

}
