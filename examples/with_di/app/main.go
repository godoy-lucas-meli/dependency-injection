package main

import (
	"fmt"
	"strings"

	"mercadolibre.com/di/examples/third_party"
	"mercadolibre.com/di/examples/with_di/repository"
	"mercadolibre.com/di/examples/with_di/service"
)

func main() {
	ra := repository.NewRActor()

	sa := service.NewSActor(ra)

	vf := third_party.NewImdbVoteFetcher()
	//vf := third_party.NewRottenTomatoesFetcher()

	sm := service.NewSMovie(sa, vf)

	name := "silvester"
	categories, err := sm.GetActorMovieCategories(name)
	if err != nil {
		panic("ups.. this should have not happened")
	}

	fmt.Printf("the actor %s has played the following movie categories: %v", name, strings.Join(categories, ", "))
}
