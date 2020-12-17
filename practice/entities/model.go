package entities

type Forecast struct {
	DateTempMap map[int64]*DailyForecast
}

type DailyForecast struct {
	MinTemp float64 `json:"min_temp,omitempty"`
	MaxTemp float64 `json:"max_temp,omitempty"`
}

type BeerPacksForecastEstimation struct {
	Date      string        `json:"timestamp,omitempty"`
	BeerPacks float64       `json:"beer_packs,omitempty"`
	Forecast  DailyForecast `json:"forecast,omitempty"`
}

type HandlerResult struct {
	Status int32
	Body   interface{}
}

type SuccessResponse struct {
	Success bool        `json:"success,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success,omitempty"`
	Message string `json:"message,omitempty"`
}

type RequestParams struct {
	Country      string `mapstructure:"country"`
	State        string `mapstructure:"state"`
	City         string `mapstructure:"city"`
	Attendees    uint   `mapstructure:"attendees"`
	ForecastDays uint   `mapstructure:"forecast_days"`
	PackUnits    uint   `mapstructure:"pack_units"`
}
