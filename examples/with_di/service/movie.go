package service

import (
	"fmt"

	domain "mercadolibre.com/di/examples/with_di"
)

type actorService interface {
	GetByActorName(name string) ([]domain.Movie, error)
}

type voteFetcher interface {
	ActorVotesByCategory(actor, category string) int
}

type SMovie struct {
	actorService actorService
	voteFetcher  voteFetcher
}

func NewSMovie(as actorService, vf voteFetcher) *SMovie {
	return &SMovie{
		actorService: as,
		voteFetcher:  vf,
	}
}

func (sm *SMovie) GetActorMovieCategories(name string) ([]string, error) {
	movies, err := sm.actorService.GetByActorName(name)
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, m := range movies {
		votes := sm.voteFetcher.ActorVotesByCategory(name, m.Category)
		categoryVotes := fmt.Sprintf("%s-%v", m.Category, votes)
		categories = append(categories, categoryVotes)
	}

	return categories, nil
}
