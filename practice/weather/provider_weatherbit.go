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
	wbURL       = internal.GetEnv("WEATHERBIT_URL", "localhost")
	wbAccessKey = internal.GetEnv("WEATHERBIT_API_KEY", "myTestAPIKey")
)

type weatherBitCli struct {
	Adapter adapter
}

func NewWeatherBitResource(adapter adapter) *weatherBitCli {
	return &weatherBitCli{
		Adapter: adapter,
	}
}

// GetAdapter gets the configured response adapter to normalize the returned data
func (cli *weatherBitCli) Adapt() adapter {
	return cli.Adapter
}

// GetForecastData fetches forecast from Weather-Bit provider
func (cli *weatherBitCli) GetForecastData(country, state, city string, forecastDays uint, httpCli httpClient) (map[string]interface{}, error) {
	URL := fmt.Sprintf("%v/v2.0/forecast/daily", wbURL)
	request := map[string]interface{}{
		"verb": http.MethodGet,
		"url":  URL,
		"queryParams": map[string]string{
			"key":     wbAccessKey,
			"country": country,
			"city":    city,
			"days":    fmt.Sprintf("%v", forecastDays),
		},
	}

	logrus.Infof("sending request to %v", URL)
	response, err := httpCli.PerformRequest(request)
	if err != nil {
		return nil, di.CustomError{
			Cause:   di.ErrDependencyNotAvailable,
			Message: "weather provider is not available",
		}
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
