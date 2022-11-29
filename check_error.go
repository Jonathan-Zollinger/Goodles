package Goodles

import (
	"os"

	"github.com/sirupsen/logrus"
)

func CheckError(err error, logger logrus.StdLogger) {
	if err != nil {
		logger.Fatal("Fatal Error:", err.Error())
		os.Exit(1)
	}
}
