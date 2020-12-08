package third_party

import (
	"math/rand"
	"time"
)

type ImdbVoteFetcher struct{}

func NewImdbVoteFetcher() *ImdbVoteFetcher {
	return &ImdbVoteFetcher{}
}

// ActorVotesByCategory fetches the amount IMDB's users votes for the current actor and category
func (ivf *ImdbVoteFetcher) ActorVotesByCategory(actor, category string) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 100
	return rand.Intn(max-min+1) + min
}
