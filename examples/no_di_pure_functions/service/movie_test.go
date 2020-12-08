package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetActorMovieCategories_success(t *testing.T) {
	categories, err := GetActorMovieCategories("silvester")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	assert.Equal(t, 2, len(categories))
	assert.ElementsMatch(t, []string{"action", "horror"}, categories)
}

func TestGetActorMovieCategories_invalidActor(t *testing.T) {
	categories, err := GetActorMovieCategories("notExistingActor")
	if err == nil {
		t.Fatalf("expected error %v", err)
	}

	assert.Nil(t, categories)
}
