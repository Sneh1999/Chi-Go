package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"https://github.com/Sneh1999/Chi-Go/todo"
)

func Routes() *chi.Mux{
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger ,  // log api request logs
		middleware.DefaultCompress, // compress results
		middleware.RedirectSlashes, //redirect slashes to o slash url versions
		middleware.Recoverer, // recover from panics thout crashing server

	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/todo",todo.Routes())
	})

	return router

}


func main(){
	router: Routes()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n",method,route)
		return nil;
	}
	if err := chi.Walk(router, walkFunc); err !=nil {
		log.Panicf("Logging err: %s\n", err.Error())

	}
	log.Fatal(http.ListenAndServe(":8080",router))
}
