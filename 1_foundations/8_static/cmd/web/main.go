package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/view", view)
	log.Println("Starting server on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
