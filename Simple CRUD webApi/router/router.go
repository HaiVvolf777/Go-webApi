package router

import (
	"github.com/gorilla/mux"
	"github.com/haider-star/mongodb/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleleOne).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleleAllMovies).Methods("DELETE")

	return router
}
