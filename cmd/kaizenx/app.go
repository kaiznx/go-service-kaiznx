package main

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kaiznx/go-service-kaiznx/endpoint"
	"github.com/kaiznx/go-service-kaiznx/service"
)

const APP_NAME = "KaizenX 0.1"

type App struct {
	router *mux.Router
	logger *logrus.Logger
}

func NewApp(logger *logrus.Logger, geocoderURL string) *App {
	// services
	geospatialService := service.NewGeospatial(geocoderURL)

	// endpoints
	geospatialEndpoint := endpoint.NewGeospatial(geospatialService)

	// routers
	r := mux.NewRouter()
	r.HandleFunc("/geocode", geospatialEndpoint.Geocode())

	r.HandleFunc("/", Index)
	return &App{r, logger}
}

func (app *App) Run(host string) {
	app.logger.Infof("Listening on %s", host)
	app.logger.Fatal(http.ListenAndServe(host, app.router))
}

func (app *App) Close() error {
	// Close here DB connections, etc
	return nil
}

func Index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s", APP_NAME)
}
