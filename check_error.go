package Goodles

import (
	"log"
	"os"
)

func CheckError(err error, log log.Logger) {
	if err != nil {
		log.Fatal("Fatal Error:", err.Error())
		os.Exit(1)
	}
}
