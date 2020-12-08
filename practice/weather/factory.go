package weather

import (
	"fmt"
)

type WeatherProvider interface {
	GetForecastData(country, state, city string, forecastDays uint, client httpClient) (map[string]interface{}, error)
	Adapt() adapter
}

// GetProvider build a new provider resource client
func GetProvider(providerName string) (WeatherProvider, error) {
	switch providerName {
	case "weather-stack":
		return NewWeatherStackResource(weatherStack), nil
	case "weather-bit":
		return NewWeatherBitResource(weatherBit), nil
	default:
		return nil, fmt.Errorf("there is no such %v defined resource to fetch the weather forecast", providerName)
	}
}
