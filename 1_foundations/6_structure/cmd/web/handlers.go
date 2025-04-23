package main

import (
	"fmt"
	"net/http"
	"strconv"
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

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed",  http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Crea algo en MyGoAPI"))
}

func view(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Muestra algo en especifico con ID %d...", id)
}
