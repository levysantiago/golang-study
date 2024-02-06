package movies_services

import movies_repository "gomovies/src/modules/movie/repositories"

func GetMoviesService() []movies_repository.Movie{
	return movies_repository.FindMany()
}