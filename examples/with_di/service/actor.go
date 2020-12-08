package service

import (
	"fmt"

	domain "mercadolibre.com/di/examples/with_di"
)

type actorRepo interface {
	FindAllMovies() []domain.Movie
}

type SActor struct {
	actorRepo actorRepo
}

func NewSActor(repo actorRepo) *SActor {
	return &SActor{
		actorRepo: repo,
	}
}

func (sa *SActor) GetByActorName(name string) ([]domain.Movie, error) {
	var movies []domain.Movie
	for _, m := range sa.actorRepo.FindAllMovies() {
		if _, ok := m.Cast[name]; ok {
			movies = append(movies, m)
		}
	}

	if len(movies) < 1 {
		return nil, fmt.Errorf("there is no movies with an actor's name such as: %s", name)
	}

	return movies, nil
}
