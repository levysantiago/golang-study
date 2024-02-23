package movies_controllers

import (
	"encoding/json"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
	"net/http"
)

type CreateMovieResponseData struct{
	Status string `json:"status"`
	Data movies_repository.Movie `json:"data"`
}

type CreateMovieResponseError struct{
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func CreateMovie(res http.ResponseWriter, req *http.Request){
	var data movies_services.CreateMovieServiceDTO

	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
			return
	}
	
	var movie = movies_services.CreateMovieService(&data)

	var responseData CreateMovieResponseData
	responseData.Status = "success"
	responseData.Data = movie

	var jsonResponse, jsonError = json.Marshal(responseData)

	if jsonError != nil {
		var responseError CreateMovieResponseError
		responseError.Status = "error"
		responseError.ErrorMessage = "Internal server error"

		jsonResponse, _ = json.Marshal(responseError)

		res.Write(jsonResponse)
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write(jsonResponse)
}