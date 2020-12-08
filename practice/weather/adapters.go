package weather

import (
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	di "mercadolibre.com/di/practice"
	"mercadolibre.com/di/practice/entities"
)

// adapter function to transform provider data into a canonical model
type adapter func(response map[string]interface{}) (*entities.Forecast, error)

var weatherStack = func(response map[string]interface{}) (*entities.Forecast, error) {
	logrus.Info("transforming weather provider weather-stack into canonical model")
	type wsForecastDaily struct {
		DateEpoch int64    `mapstructure:"date_epoch"`
		MinTemp   *float64 `mapstructure:"mintemp"`
		MaxTemp   *float64 `mapstructure:"maxtemp"`
	}
	type wsForecast struct {
		Forecast map[string]*wsForecastDaily `mapstructure:"forecast"`
	}

	var wsf wsForecast
	if err := mapstructure.Decode(response, &wsf); err != nil {
		return nil, err
	}

	resultMap := make(map[int64]*entities.DailyForecast, len(wsf.Forecast))
	for _, day := range wsf.Forecast {
		if day.MaxTemp == nil && day.MinTemp == nil {
			return nil, di.ErrResourceMissingData
		}

		df := &entities.DailyForecast{}
		if day.MaxTemp != nil {
			df.MaxTemp = *day.MaxTemp
		}

		if day.MinTemp != nil {
			df.MinTemp = *day.MinTemp
		}

		resultMap[day.DateEpoch] = df
	}

	return &entities.Forecast{DateTempMap: resultMap}, nil
}

var weatherBit = func(response map[string]interface{}) (*entities.Forecast, error) {
	logrus.Info("transforming weather provider weather-bit into canonical model")
	type wbForecastDaily struct {
		Datetime *string  `mapstructure:"valid_date"`
		MinTemp  *float64 `mapstructure:"min_temp"`
		MaxTemp  *float64 `mapstructure:"max_temp"`
	}
	type wbForecast struct {
		Forecast []*wbForecastDaily `mapstructure:"data"`
	}

	var wbf wbForecast
	if err := mapstructure.Decode(response, &wbf); err != nil {
		return nil, err
	}

	resultMap := make(map[int64]*entities.DailyForecast, len(wbf.Forecast))
	for _, day := range wbf.Forecast {
		if day.MaxTemp == nil || day.Datetime == nil {
			return nil, di.ErrResourceMissingData
		}

		actualDate, err := time.Parse("2006-01-02", *day.Datetime)
		if err != nil {
			return nil, err
		}
		ts := actualDate.Unix()

		df := &entities.DailyForecast{}
		if day.MaxTemp != nil {
			df.MaxTemp = *day.MaxTemp
		}

		if day.MinTemp != nil {
			df.MinTemp = *day.MinTemp
		}

		resultMap[ts] = df
	}

	return &entities.Forecast{DateTempMap: resultMap}, nil
}
