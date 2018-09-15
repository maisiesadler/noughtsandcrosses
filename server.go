package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/maisiesadler/noughtsandcrosses/game"
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

func main() {
	go startServer()

	// quit on new line
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", game.String())
	})

	http.HandleFunc("/place", func(w http.ResponseWriter, r *http.Request) {
		var request placeRequest
		readBody(w, r, &request)
		err := game.PlaceCounter(request.Counter, request.X, request.Y)
		if err != nil {
			log.Printf("Error placing counter: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, game.String())
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		game.Reset()
		fmt.Fprintf(w, "ok")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
