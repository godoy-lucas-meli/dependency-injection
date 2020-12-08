package business

import (
	"github.com/sirupsen/logrus"
	"mercadolibre.com/di/practice/entities"
)

type forecastFetcher interface {
	Get(country, state, city string, forecastDays uint) (*entities.Forecast, error)
}

type BeerForecast struct {
	forecastFetcher forecastFetcher
}

func NewBeerForecast(forecastFetcher forecastFetcher) *BeerForecast {
	return &BeerForecast{forecastFetcher: forecastFetcher}
}

func (bf *BeerForecast) Get(rp *entities.RequestParams) ([]*entities.BeerForecast, error) {
	forecast, err := bf.getForecast(rp.Country, rp.State, rp.City, rp.ForecastDays)
	if err != nil {
		return nil, err
	}

	var results []*entities.BeerForecast
	for ts, f := range forecast.DateTempMap {
		bp, err := beerPacksQuantity(rp.Attendees, rp.PackUnits, f.MaxTemp)
		if err != nil {
			return nil, err
		}

		results = append(results, &entities.BeerForecast{
			Timestamp: &ts,
			BeerPacks: &bp,
			Forecast: &entities.DailyForecast{
				MinTemp: f.MinTemp,
				MaxTemp: f.MaxTemp,
			},
		})
	}

	return results, nil
}

func (bf *BeerForecast) getForecast(country, city, state string, forecastDays uint) (*entities.Forecast, error) {
	forecast, err := bf.forecastFetcher.Get(country, state, city, forecastDays)
	if err != nil {
		return nil, err
	}

	logrus.Infof("forecast values for %v, %v, %v are: %v", country, state, city, forecast)
	return forecast, nil
}
