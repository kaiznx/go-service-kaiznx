package endpoint

import (
	"net/http"

	"github.com/kaiznx/go-service-kaiznx/service"
)

type (
	Geospatial interface {
		Geocode() Endpoint
	}

	geospatialEndpoint struct {
		geospatialService service.Geospatial
	}
)

func NewGeospatial(geospatialService service.Geospatial) Geospatial {
	return &geospatialEndpoint{geospatialService}
}

func (e geospatialEndpoint) Geocode() Endpoint {
	return func(w http.ResponseWriter, req *http.Request) {
		// Get query values
		queryValues := req.URL.Query()
		address := queryValues.Get("q")

		resp, err := e.geospatialService.Geocode(address)
		if err != nil {
			http.Error(w, err.Error(), 400)
		}
		JSON(w, resp)
	}
}
