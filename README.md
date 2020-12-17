## DOJO - Dependency Injection - The Go way

#### Files Content
```
dependency-injection
    examples                    (trivial examples)
        no_di_pure_functions    (using functions with no dependency injection)
        third_party             (package to represent some 3rd party library)
        with_di                 (using dependency injection)
    
    practice                    (beer forecast project to exercise the dojo's content)
```

#### Configuration
##### Env Vars
In order to run, you need to fill these env vars. For API keys provide you may need to get a fresh token by creating an
account on the weather provider's sites.

```
WEATHERBIT_API_KEY=<your-key>
WEATHERSTACK_API_KEY=<your-key>
WEATHER_PROVIDER=weather-bit
WEATHERSTACK_URL=http://api.weatherstack.com
WEATHERBIT_URL=https://api.weatherbit.io
```

### Practice Time
In order to improve the strong coupling and assign the right responsibility to each app layer, refactor all the
components between the handler controller and the weather forecast client using dependency injection pattern.

#### Beer Forecast
This API exposes a `GET` endpoint which retrieves the amount of beer packs you need to buy for an event based on the
amount of people attending it and the weather forecast.

##### Rules
How many beer cans/bottles an attendant could drink based on the current weather based on the temperature:
```
if temperature >= 20 && temperature <= 24 {
		return 2
	} else if temperature < 20 {
		return 1
	}
return 3 
```

- The very first aim of the exercise, it is to reach the point where you can replace all the package coupling by using `interface`s.
- The second goal, would be to unit test the method `Estimate` from the `estimator.go` file, by mocking its dependencies.

#### Hitting the service

```
curl --location --request GET \
    'http://localhost:3001/beer-forecast?country=argentina&state=cordoba&city=cordoba&attendees=10&forecast_days=10&pack_units=2'
```

Query Parameters
```
country:        the country name where you want to run the event
state:          the state name where you want to run the event
city:           the city name where you want to run the event

attendees:      the number of attendants the events will host
forecast_days:  amount of days ahead you want to analyze the forecast ahead (max 10)
pack_units:     beer quantity per pack     

```

##### Response example
```
{
    "success": true,
    "data": [
        {
            "timestamp": 1607904000,
            "beer_packs": 15,
            "forecast": {
                "min_temp": 15.5,
                "max_temp": 32
            }
        },
        {
            "timestamp": 1608249600,
            "beer_packs": 15,
            "forecast": {
                "min_temp": 16,
                "max_temp": 33.5
            }
        },
        ...
    ]
}

```
