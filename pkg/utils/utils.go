package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ParseBody(r *http.Request, x any) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	log.Printf("Received body: %s", body)

	if err := json.Unmarshal(body, x); err != nil {
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}
}
