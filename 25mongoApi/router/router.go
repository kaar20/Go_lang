package router

import (
	"github.com/gorilla/mux"
	"github.com/kaar20/mongoapi/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMoviesList).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	router.HandleFunc("/api/deleteMovie/{id}", controller.DeleteOne).Methods("DELETE")

	router.HandleFunc("/api/deleteMovies", controller.DeleteAll).Methods("DELETE")

	return router
}
