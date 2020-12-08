package third_party

import (
	"math/rand"
	"time"
)

type RottenTomatoesVoteFetcher struct{}

func NewRottenTomatoesFetcher() *RottenTomatoesVoteFetcher {
	return &RottenTomatoesVoteFetcher{}
}

// ActorVotesByCategory fetches the amount Rotten Tomattoes' users votes for the current actor and category
func (ivf *RottenTomatoesVoteFetcher) ActorVotesByCategory(actor, category string) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 5
	return rand.Intn(max-min+1) + min
}
