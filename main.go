package main

import (
	"github.com/lucianorbr/formWeb_Go/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
