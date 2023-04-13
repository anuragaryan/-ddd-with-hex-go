package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func responseJSON(w http.ResponseWriter, statusCode int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if resp == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("Unable to send response to client: %v", err)
	}
}
