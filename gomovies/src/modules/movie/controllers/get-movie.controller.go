package movies_controllers

import (
	"encoding/json"
	"fmt"
	movies_services "gomovies/src/modules/movie/services"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMovie(res http.ResponseWriter, req *http.Request){
	var id = mux.Vars(req)["id"]

	var movie = movies_services.GetMovie(id)

	jsonResponse, jsonError := json.Marshal(movie)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	res.WriteHeader(http.StatusOK)
	res.Write(jsonResponse)
}