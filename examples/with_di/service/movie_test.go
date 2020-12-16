package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	domain "mercadolibre.com/di/examples/with_di"
)

type actorServiceMock struct {
	getByActorNameMock func(name string) ([]domain.Movie, error)
}

func (a actorServiceMock) GetByActorName(name string) ([]domain.Movie, error) {
	return a.getByActorNameMock(name)
}

type voteFetcherMock struct {
	actorVotesByCategoryMock func(actor, category string) int
}

func (v voteFetcherMock) ActorVotesByCategory(actor, category string) int {
	return v.actorVotesByCategoryMock(actor, category)
}

func TestSMovie_GetActorMovieCategories(t *testing.T) {
	asMock := actorServiceMock{func(name string) ([]domain.Movie, error) {
		return []domain.Movie{
			{
				Name:     "movie 1",
				Category: "horror",
			},
			{
				Name:     "movie 2",
				Category: "drama",
			},
		}, nil
	}}

	vfMock := voteFetcherMock{func(actor, category string) int {
		if category == "horror" {
			return 10
		}
		return 5
	}}

	expected := []string{"horror-10", "drama-5"}

	movieService := NewSMovie(asMock, vfMock)

	result, err := movieService.GetActorMovieCategories("someOne")

	assert.Nil(t, err)
	assert.NotEmpty(t, result)
	assert.ElementsMatch(t, expected, result)
}

func TestSMovie_GetActorMovieCategories_ErrorFromActorService(t *testing.T) {
	errMsg := "something went wrong here"

	asMock := actorServiceMock{func(name string) ([]domain.Movie, error) {
		return nil, errors.New(errMsg)
	}}

	vfMock := voteFetcherMock{func(actor, category string) int {
		return 0
	}}

	movieService := NewSMovie(asMock, vfMock)

	result, err := movieService.GetActorMovieCategories("someOne")

	assert.NotNil(t, err)
	assert.Equal(t, errMsg, err.Error())
	assert.Empty(t, result)
}
