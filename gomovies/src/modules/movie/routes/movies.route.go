package routes

import (
	movies_controllers "gomovies/src/modules/movie/controllers"

	"github.com/gorilla/mux"
)

func MoviesRoutes(router *mux.Router){
	// Get Movies
	router.HandleFunc("/movies", movies_controllers.GetMovies).Methods("GET")
	// Get a Movie
	router.HandleFunc("/movies/{id}", movies_controllers.GetMovie).Methods("GET")
	// Get Create Movie
	router.HandleFunc("/movies", movies_controllers.CreateMovie).Methods("POST")
	// Get Update movie
	router.HandleFunc("/movies/{id}", movies_controllers.UpdateMovie).Methods("PUT")
	// Get Delete Movie
	router.HandleFunc("/movies/{id}", movies_controllers.DeleteMovie).Methods("DELETE")
}