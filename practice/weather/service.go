package weather

import (
	"github.com/sirupsen/logrus"
	"mercadolibre.com/di/practice/entities"
	"mercadolibre.com/di/practice/internal"
)

const ttl = 30

var weatherProvider = internal.GetEnv("WEATHER_PROVIDER", "weather-bit")

type Service struct {
	weatherProvider WeatherProvider
}

// NewWeatherService creates new weather Service instance
func NewWeatherService() (*Service, error) {
	wProvider, err := GetProvider(weatherProvider)
	if err != nil {
		panic(err)
	}

	return &Service{weatherProvider: wProvider}, nil
}

// Get gets forecast based on configured weather provider
func (s *Service) Get(country, state, city string, forecastDays uint) (*entities.Forecast, error) {
	logrus.Info("fetching forecast from weather provider")

	client := internal.NewHttpClient(ttl)

	forecastData, err := s.weatherProvider.GetForecastData(country, state, city, forecastDays, client)
	if err != nil {
		return nil, err
	}

	forecast, err := s.weatherProvider.Adapt()(forecastData)
	if err != nil {
		return nil, err
	}

	return forecast, nil
}
