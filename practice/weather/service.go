package weather

import (
	"github.com/sirupsen/logrus"
	"mercadolibre.com/di/practice/entities"
	"mercadolibre.com/di/practice/internal"
)

const ttl = 30

type FService struct {
	weatherProvider WeatherProvider
}

// NewForecastService creates new weather service instance
func NewForecastService(weatherProvider WeatherProvider) (*FService, error) {
	return &FService{weatherProvider: weatherProvider}, nil
}

// Get gets forecast based on configured weather provider
func (ws *FService) Get(country, state, city string, forecastDays uint) (*entities.Forecast, error) {
	logrus.Info("fetching forecast from weather provider")
	client := internal.NewHttpClient(ttl)
	forecastData, err := ws.weatherProvider.GetForecastData(country, state, city, forecastDays, client)
	if err != nil {
		return nil, err
	}

	forecast, err := ws.weatherProvider.Adapt()(forecastData)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}
