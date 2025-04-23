package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola desde Mi Primera API con Go"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Println("Starting server on http://localhost:9876")
	err := http.ListenAndServe(":9876", mux)
	log.Fatal(err)
}
