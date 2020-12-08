package service

import (
	"fmt"

	"mercadolibre.com/di/examples/no_di_pure_functions"
	"mercadolibre.com/di/examples/no_di_pure_functions/repository"
)

func getByActorName(name string) ([]domain.Movie, error) {
	var movies []domain.Movie
	for _, m := range repository.FindAllMovies() {
		if _, ok := m.Cast[name]; ok {
			movies = append(movies, m)
		}
	}

	if len(movies) < 1 {
		return nil, fmt.Errorf("there is no movies with an actor's name such as: %s", name)
	}

	return movies, nil
}
