package movies_services

import movies_repository "gomovies/src/modules/movie/repositories"

func GetMovie(id string) movies_repository.Movie{
	return movies_repository.FindById(id)
}