package main

import (
	"fmt"
	"strings"

	"mercadolibre.com/di/examples/no_di_pure_functions/service"
)

func main() {
	name := "silvester"
	categories, err := service.GetActorMovieCategories(name)
	if err != nil {
		panic("ups.. this should have not happened")
	}

	fmt.Printf("the actor %s has played the following movie categories: %v", name, strings.Join(categories, ", "))
}
