package movies_repository

import "fmt"

type Movie struct{
	Id string `json:"Id"`
	Isbn string `json"Isbn"`
	Title string `json"Title"`
	Director_Id string `json"Director_Id"`
}

var movies []Movie

func Create(newMovie Movie) {
	movies = append(movies, newMovie)
	
	fmt.Println(movies)
}

func FindById(id string) Movie{
	var foundMovie Movie
	
	for _, m := range movies {
		if(m.Id == id){
			foundMovie = m
		}
	}

	return foundMovie
}