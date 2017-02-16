package model

import (
	"github.com/paulmach/go.geojson"
)

//
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []*Feature `json:"features"`
}

//
type Feature struct {
	*geojson.Feature
}
