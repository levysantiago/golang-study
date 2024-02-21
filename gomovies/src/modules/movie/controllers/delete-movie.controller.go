package movies_controllers

import (
	"encoding/json"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
	"net/http"

	"github.com/gorilla/mux"
)

type DeleteMovieResponseError struct{
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

type DeleteMovieResponseData struct{
	Status string `json:"status"`
	Data movies_repository.Movie `json:"data"`
}

func DeleteMovie(res http.ResponseWriter, req *http.Request){
	var movieId = mux.Vars(req)["id"]

	var deleteMovieData = movies_services.DeleteMovieServiceDTO{MovieId: movieId}

	movie, err := movies_services.DeleteMovie(&deleteMovieData)

	if(err != nil) {
		var responseError DeleteMovieResponseError
		responseError.Status = "error"
		responseError.ErrorMessage = err.Error()

		jsonResponse, _ := json.Marshal(responseError)

		res.WriteHeader(http.StatusBadRequest)
		res.Write(jsonResponse)
		return 
	}

	var responseData DeleteMovieResponseData
	responseData.Status = "success"
	responseData.Data = movie

	jsonResponse, jsonError := json.Marshal(responseData)

	if(jsonError != nil){
		var responseError DeleteMovieResponseError
		responseError.Status = "error"
		responseError.ErrorMessage = "Internal Server Error"

		res.WriteHeader(http.StatusInternalServerError)
		res.Write(jsonResponse)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponse)
}