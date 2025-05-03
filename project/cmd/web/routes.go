package main

import (
	"net/http"
	"github.com/justinas/alice"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer))
	dynamic := alice.New(app.sessionManager.LoadAndSave)
	router.Handler(http.MethodGet, "/",  dynamic.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/view/:id",  dynamic.ThenFunc(app.view))
	router.Handler(http.MethodGet, "/create", dynamic.ThenFunc(app.create))
	router.Handler(http.MethodPost, "/create", dynamic.ThenFunc(app.createPost))
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
