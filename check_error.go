package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func checkError(err error, logger logrus.StdLogger) {
	if err != nil {
		logger.Fatal("Fatal Error:", err.Error())
		os.Exit(1)
	}
}
