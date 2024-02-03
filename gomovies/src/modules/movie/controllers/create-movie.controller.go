package movies_controllers

import (
	"encoding/json"
	movies_services "gomovies/src/modules/movie/services"
	"log"
	"net/http"
)



func CreateMovie(res http.ResponseWriter, req *http.Request){
	var data movies_services.CreateMovieServiceDTO

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
	}
	
	var movie = movies_services.CreateMovieService(&data)

	var jsonResponse, jsonError = json.Marshal(movie)

	if jsonError != nil {
		log.Fatal("Unable to encode JSON")
	}

	res.WriteHeader(http.StatusCreated)
	res.Write(jsonResponse)
}