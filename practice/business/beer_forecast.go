package business

import (
	"github.com/sirupsen/logrus"
	"mercadolibre.com/di/practice/entities"
)

type weatherFetcher interface {
	Get(country, state, city string, forecastDays uint) (*entities.Forecast, error)
}

type estimator struct {
	wFetcher weatherFetcher
}

func NewBeerPacksEstimator(wf weatherFetcher) *estimator {
	return &estimator{wFetcher: wf}
}

func (e *estimator) Estimate(rp *entities.RequestParams) ([]*entities.BeerPacksForecastEstimation, error) {
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

		results = append(results, &entities.BeerPacksForecastEstimation{
			Timestamp: &timestamp,
			BeerPacks: &qty,
			Forecast: &entities.DailyForecast{
				MinTemp: df.MinTemp,
				MaxTemp: df.MaxTemp,
			},
		})
	}

	return results, nil
}

func (e *estimator) getForecast(country, city, state string, forecastDays uint) (*entities.Forecast, error) {
	forecast, err := e.wFetcher.Get(country, state, city, forecastDays)
	if err != nil {
		return nil, err
	}

	logrus.Infof("forecast values for %v, %v, %v are: %v", country, state, city, forecast)
	return forecast, nil
}
