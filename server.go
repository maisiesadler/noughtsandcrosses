package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maisiesadler/noughtsandcrosses/game"
)

func main() {
	startServer()
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
