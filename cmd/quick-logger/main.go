package main

import (
	"fmt"

	"github.com/gbburleigh/quick-logger/pkg/logger" // Replace yourmodule
)

func main() {
	log := logger.NewLogger()

	log.Info("Application started")
	log.Warn("Low disk space")
	log.Error("Database connection failed")
	log.Critical("System is down")

	fmt.Println("Check console for logs")
}
