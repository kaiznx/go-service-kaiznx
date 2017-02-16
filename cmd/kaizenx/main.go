package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	port := os.Getenv("PORT")
	if port == "" {
		logger.WithField("PORT", port).Fatal("$PORT must be set")
	}

	geocoderURL := os.Getenv("GEOCODER_URL")
	if geocoderURL == "" {
		geocoderURL = "http://search.mapzen.com/v1/"
	}

	app := NewApp(logger, geocoderURL)

	signalHandler(app)

	app.Run(":" + port)
}

func signalHandler(app *App) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	// Block until a signal is received
	go func() {
		sig := <-c
		if err := app.Close(); err != nil {
			app.logger.Errorf("Failed to stop container manager: %v", err)
		}
		app.logger.Infof("Exiting given signal: %v", sig)
		os.Exit(0)
	}()
}
