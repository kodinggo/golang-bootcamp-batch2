package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if os.Getenv("env") == "development" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// logrus.Fatal("failed connect db")
	// logrus.Panic("failed connect db")

	logrus.WithFields(logrus.Fields{
		"product_id": 1,
		"user_id": 1,
	}).Error("someting wrong")

	logrus.Warn("credential almost expired")
	logrus.Info("payment success with ref_id xxxx")
	logrus.Debug("debugging something")
}
