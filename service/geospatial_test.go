package service

import (
	"os"
	"testing"
)

func TestGeocoding(t *testing.T) {
	geocoderURL := os.Getenv("GEOCODER_URL")
	if geocoderURL == "" {
		geocoderURL = "http://search.mapzen.com/v1/"
	}

	geospatialService := NewGeospatial(geocoderURL)

	addrTests := []struct {
		input    string
		expected int
	}{
		{
			input:    "4868 W Flagler ST Miami FL",
			expected: 1,
		},
		{
			input:    "701 NW 62nd Ave #400, Miami, FL 33126",
			expected: 1,
		},
		{
			input:    "9555 SW 162nd Ave Miami FL 33196",
			expected: 1,
		},
	}

	for i, c := range addrTests {
		feature, err := geospatialService.Geocode(c.input)
		if err != nil {
			t.Errorf("case: %d input: %v error: %v", i+1, c.input, err)
			continue
		}
		if feature == nil || len(feature.Features) != c.expected {
			t.Errorf("case: %d input: %v \n %v", i+1, c.input, feature)
		}
	}
}
