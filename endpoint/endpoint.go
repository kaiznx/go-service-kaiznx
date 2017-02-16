package endpoint

import (
	"encoding/json"
	"io"
	"net/http"
)

type Endpoint func(http.ResponseWriter, *http.Request)

func JSON(w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func DecodeReq(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(&v)
}
