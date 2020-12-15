package business

import (
	"errors"
	"math"
)

// beerPacksQuantity calculates the amount of beer packs based on number of attendees, temperature and how many beers
// are contained per pack
func beerPacksQuantity(attendees, unitsPerPack uint, temperature float64) (float64, error) {
	if unitsPerPack < 1 {
		return 0, errors.New("min amount of bottle/cans per pack has to be a positive number")
	}

	if attendees == 0 {
		return 0, nil
	}

	multiplier := getMultiplier(temperature)

	result := math.Ceil((float64(attendees)*multiplier/float64(unitsPerPack))*100) / 100

	return result, nil
}

func getMultiplier(temperature float64) float64 {
	if temperature >= 20 && temperature <= 24 {
		return 2
	} else if temperature < 20 {
		return 1
	}
	return 3
}
