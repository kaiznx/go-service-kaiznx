package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kaiznx/go-service-kaiznx/model"
)

// Geospatial ...
type (
	Geospatial interface {
		Geocode(address string) (*model.FeatureCollection, error)
	}

	geospatialService struct {
		baseURL string
	}
)

// NewGeospatial creates a geospatial service
func NewGeospatial(baseURL string) Geospatial {
	return &geospatialService{baseURL}
}

func (s geospatialService) Geocode(address string) (*model.FeatureCollection, error) {
	uri := s.baseURL + "search?size=1&text=%s"

	resp, err := http.Get(fmt.Sprintf(uri, url.QueryEscape(address)))
	if err != nil {
		return nil, err
	}

	var feature *model.FeatureCollection
	if err := json.NewDecoder(resp.Body).Decode(&feature); err != nil {
		return nil, err
	}

	return feature, nil
}
