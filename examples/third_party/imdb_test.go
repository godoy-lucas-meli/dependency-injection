package third_party

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestActorVotesByCategory(t *testing.T) {
	for i := 0; i <= 10; i++ {
		votesByCategory := ActorVotesByCategory("a", "b")
		assert.Greater(t, votesByCategory, 0)
		assert.Less(t, votesByCategory, 101)
	}
}
