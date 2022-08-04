package main

import (
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/documents", Document)
	http.HandleFunc("/jsondocuments", JsonDocument)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
