package movies_repository

import "fmt"

type Movie struct{
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director_Id string `json:"director_id"`
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