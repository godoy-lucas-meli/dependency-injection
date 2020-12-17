package business

import (
	"github.com/sirupsen/logrus"
	"mercadolibre.com/di/practice/entities"
	"mercadolibre.com/di/practice/weather"
)

type Estimator struct {
	wService *weather.Service
}

func NewBeerPacksEstimator() *Estimator {
	weatherService, err := weather.NewWeatherService()
	if err != nil {
		panic(err)
	}

	return &Estimator{wService: weatherService}
}

func (e *Estimator) Estimate(rp *entities.RequestParams) ([]*entities.BeerPacksForecastEstimation, error) {
	forecast, err := e.getForecast(rp.Country, rp.State, rp.City, rp.ForecastDays)
	if err != nil {
		return nil, err
	}

	var results []*entities.BeerPacksForecastEstimation
	for timestamp, df := range forecast.DateTempMap {
		qty, err := beerPacksQuantity(rp.Attendees, rp.PackUnits, df.MaxTemp)
		if err != nil {
			return nil, err
		}

		ts := timestamp
		results = append(results, &entities.BeerPacksForecastEstimation{
			Timestamp: ts,
			BeerPacks: qty,
			Forecast: &entities.DailyForecast{
				MinTemp: df.MinTemp,
				MaxTemp: df.MaxTemp,
			},
		})
	}

	return results, nil
}

func (e *Estimator) getForecast(country, city, state string, forecastDays uint) (*entities.Forecast, error) {
	forecast, err := e.wService.Get(country, state, city, forecastDays)
	if err != nil {
		return nil, err
	}

	logrus.Infof("forecast values for %v, %v, %v are: %v", country, state, city, forecast)
	return forecast, nil
}
