package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	di "mercadolibre.com/di/practice"
	"mercadolibre.com/di/practice/internal"
)

var (
	wsURL       = internal.GetEnv("WEATHERSTACK_URL", "test.com")
	wsAccessKey = internal.GetEnv("WEATHERSTACK_API_KEY", "myTestAPIKey")
)

type weatherStackCli struct {
	Adapter adapter
}

func NewWeatherStackResource(adapter adapter) *weatherStackCli {
	return &weatherStackCli{
		Adapter: adapter,
	}
}

// GetAdapter gets the configured response adapter to normalize the returned data
func (cli *weatherStackCli) Adapt() adapter {
	return cli.Adapter
}

// GetForecastData fetches forecast from Weather-Stack provider
func (cli *weatherStackCli) GetForecastData(country, state, city string, forecastDays uint, httpCli httpClient) (map[string]interface{}, error) {
	URL := fmt.Sprintf("%v/forecast", wsURL)
	request := map[string]interface{}{
		"verb": http.MethodGet,
		"url":  URL,
		"queryParams": map[string]string{
			"access_key":    wsAccessKey,
			"query":         fmt.Sprintf("%v,%v,%v", country, state, city),
			"forecast_days": fmt.Sprintf("%v", forecastDays),
			"units":         "m",
		},
	}

	logrus.Infof("sending request to %v", URL)
	response, err := httpCli.PerformRequest(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, di.CustomError{
			Cause:   di.ErrDependencyNotAvailable,
			Message: fmt.Sprintf("weather provider responded with an invalid status code: %d", response.StatusCode),
		}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var bodyMap map[string]interface{}

	if err := json.Unmarshal(body, &bodyMap); err != nil {
		return nil, err
	}

	return bodyMap, nil
}
