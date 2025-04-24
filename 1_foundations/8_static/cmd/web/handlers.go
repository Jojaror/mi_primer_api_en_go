package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func download(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/static/file.zip")
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
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
