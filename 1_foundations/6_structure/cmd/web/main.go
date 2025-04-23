package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/view", view)
	log.Println("Starting server on http://localhost:9876")
	err := http.ListenAndServe(":9876", mux)
	log.Fatal(err)
}
