package business

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"mercadolibre.com/di/practice/entities"
)

type weatherFetcherMock struct {
	GetMock func(country, state, city string, forecastDays uint) (*entities.Forecast, error)
}

func (wfm weatherFetcherMock) Get(country, state, city string, forecastDays uint) (*entities.Forecast, error) {
	return wfm.GetMock(country, state, city, forecastDays)
}

func TestEstimator_Estimate(t *testing.T) {
	wfMock := &weatherFetcherMock{
		GetMock: func(country, state, city string, forecastDays uint) (*entities.Forecast, error) {
			return &entities.Forecast{DateTempMap: map[int64]*entities.DailyForecast{
				1608595200: {
					MinTemp: 10,
					MaxTemp: 19,
				},
				1608681600: {
					MinTemp: 10,
					MaxTemp: 39,
				},
			}}, nil
		}}

	estimator := NewBeerPacksEstimator(wfMock)

	expected := []*entities.BeerPacksForecastEstimation{
		{
			Timestamp: int64Pointer(1608595200),
			BeerPacks: float64Pointer(1),
			Forecast: &entities.DailyForecast{
				MinTemp: 10,
				MaxTemp: 19,
			},
		},
		{
			Timestamp: int64Pointer(1608681600),
			BeerPacks: float64Pointer(3),
			Forecast: &entities.DailyForecast{
				MinTemp: 10,
				MaxTemp: 39,
			},
		},
	}

	estimate, err := estimator.Estimate(&entities.RequestParams{
		Country:      "argentina",
		State:        "cordoba",
		City:         "cordoba",
		Attendees:    6,
		ForecastDays: 2,
		PackUnits:    6,
	})

	if err != nil {
		t.Fatalf("unexpected error with value: %v", err)
	}

	assert.EqualValues(t, expected, estimate)
}

func TestEstimator_Estimate_ErrorFetchingWeather(t *testing.T) {
	wfMock := &weatherFetcherMock{
		GetMock: func(country, state, city string, forecastDays uint) (*entities.Forecast, error) {
			return nil, errors.New("some error occurred")
		}}

	estimator := NewBeerPacksEstimator(wfMock)

	estimate, err := estimator.Estimate(&entities.RequestParams{
		Country:      "argentina",
		State:        "cordoba",
		City:         "cordoba",
		Attendees:    6,
		ForecastDays: 2,
		PackUnits:    6,
	})

	assert.Nil(t, estimate)
	assert.NotNil(t, err)
}

func float64Pointer(n float64) *float64 {
	return &n
}

func int64Pointer(n int64) *int64 {
	return &n
}
