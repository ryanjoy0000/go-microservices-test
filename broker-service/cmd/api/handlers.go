package main

import (
	"encoding/json"
	"net/http"
)

type responseJson struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
}

func (conf *BrokerAppConfig) BrokerPost(w http.ResponseWriter, r *http.Request) {

	// Create a POST response (to return to the client's request)
	resp := responseJson{
		Message: "You reached Broker Service ...",
		Error:   false,
	}

	// Convert json to byte slice
	bSlice, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(bSlice)
}
