package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func readBody(w http.ResponseWriter, r *http.Request, o interface{}) interface{} {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return nil
	}

	// Unmarshal
	err = json.Unmarshal(body, &o)
	if err != nil {
		log.Printf("Error parsing json: %v", err)
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return nil
	}

	return o
}
