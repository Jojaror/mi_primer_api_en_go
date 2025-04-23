package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from MyGoAPI"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Read about MyGoAPI"))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Conctact the administrator of MyGoAPI"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	log.Println("Starting server on http://localhost:9876")
	err := http.ListenAndServe(":9876", mux)
	log.Fatal(err)
}
