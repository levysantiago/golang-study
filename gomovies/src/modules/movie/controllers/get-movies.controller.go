package movies_controllers

import (
	"encoding/json"
	movies_repository "gomovies/src/modules/movie/repositories"
	movies_services "gomovies/src/modules/movie/services"
	"net/http"
)

type GetMoviesResponseData struct{
	Status string `json:"status"`
	Data []movies_repository.Movie `json:"data"`
}

type GetMoviesResponseError struct{
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func GetMovies(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")

	var movies = movies_services.GetMoviesService()

	var responseData GetMoviesResponseData
	if len(movies) > 0 {
		responseData.Status = "success"
		responseData.Data = movies
	} else {
		responseData = GetMoviesResponseData{Data: []movies_repository.Movie{}, Status: "success"}
	}

	jsonResponse, jsonError := json.Marshal(responseData)

	if jsonError != nil {
		var responseError GetMoviesResponseError
		responseError.Status = "error"
		responseError.ErrorMessage = "Internal server error"

		jsonResponse, _ = json.Marshal(responseError)
		
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(jsonResponse)
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponse)
}