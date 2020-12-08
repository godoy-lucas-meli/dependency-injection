package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mercadolibre/go-meli-toolkit/restful/rest"
	"gopkg.in/go-playground/assert.v1"
)

const weatherBitMockResponse = `{"data":[{"moonrise_ts":1607405745,"wind_cdir":"NE","rh":46,"pres":955.241,"high_temp":31.3,"sunset_ts":1607469602,"ozone":280.777,"moon_phase":0.309648,"wind_gust_spd":7.7312,"snow_depth":0,"clouds":0,"ts":1607396460,"sunrise_ts":1607418320,"app_min_temp":18,"wind_spd":4.03654,"pop":0,"wind_cdir_full":"northeast","slp":1013.05,"moon_phase_lunation":0.83,"valid_date":"2020-12-08","app_max_temp":30.8,"vis":24.1349,"dewpt":12.9,"snow":0,"uv":11.847,"weather":{"icon":"c01d","code":800,"description":"Clear Sky"},"wind_dir":42,"max_dhi":null,"clouds_hi":0,"precip":0,"low_temp":21.2,"max_temp":31.4,"moonset_ts":1607449766,"datetime":"2020-12-08","temp":25.7,"min_temp":18.3,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1607494152,"wind_cdir":"ESE","rh":41,"pres":955.02,"high_temp":34.2,"sunset_ts":1607556046,"ozone":265.004,"moon_phase":0.202744,"wind_gust_spd":7.81936,"snow_depth":0,"clouds":0,"ts":1607482860,"sunrise_ts":1607504731,"app_min_temp":21,"wind_spd":3.58692,"pop":0,"wind_cdir_full":"east-southeast","slp":1012.5,"moon_phase_lunation":0.86,"valid_date":"2020-12-09","app_max_temp":33.1,"vis":24.135,"dewpt":12.6,"snow":0,"uv":11.9157,"weather":{"icon":"c01d","code":800,"description":"Clear Sky"},"wind_dir":119,"max_dhi":null,"clouds_hi":0,"precip":0,"low_temp":22.3,"max_temp":34.5,"moonset_ts":1607540188,"datetime":"2020-12-09","temp":27.6,"min_temp":21.1,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1607582652,"wind_cdir":"ENE","rh":35,"pres":946.722,"high_temp":34.5,"sunset_ts":1607642488,"ozone":248.883,"moon_phase":0.112274,"wind_gust_spd":14.3,"snow_depth":0,"clouds":2,"ts":1607569260,"sunrise_ts":1607591143,"app_min_temp":21.8,"wind_spd":6.7674,"pop":0,"wind_cdir_full":"east-northeast","slp":1004.26,"moon_phase_lunation":0.9,"valid_date":"2020-12-10","app_max_temp":33,"vis":24.135,"dewpt":11,"snow":0,"uv":11.9797,"weather":{"icon":"c02d","code":801,"description":"Few clouds"},"wind_dir":61,"max_dhi":null,"clouds_hi":2,"precip":0,"low_temp":15.3,"max_temp":34.5,"moonset_ts":1607630737,"datetime":"2020-12-10","temp":28.3,"min_temp":22.4,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1607671346,"wind_cdir":"SSE","rh":38,"pres":951.314,"high_temp":26.5,"sunset_ts":1607728930,"ozone":262.631,"moon_phase":0.0454814,"wind_gust_spd":19.2067,"snow_depth":0,"clouds":15,"ts":1607655660,"sunrise_ts":1607677556,"app_min_temp":15.3,"wind_spd":5.74468,"pop":20,"wind_cdir_full":"south-southeast","slp":1007.57,"moon_phase_lunation":0.93,"valid_date":"2020-12-11","app_max_temp":25.8,"vis":24.1351,"dewpt":5.1,"snow":0,"uv":11.9019,"weather":{"icon":"c02d","code":801,"description":"Few clouds"},"wind_dir":151,"max_dhi":null,"clouds_hi":0,"precip":0.0625,"low_temp":15.7,"max_temp":29.4,"moonset_ts":1607721412,"datetime":"2020-12-11","temp":21.8,"min_temp":15.2,"clouds_mid":0,"clouds_low":15},{"moonrise_ts":1607760335,"wind_cdir":"SSE","rh":23,"pres":950.989,"high_temp":31.6,"sunset_ts":1607815370,"ozone":269.802,"moon_phase":0.00773562,"wind_gust_spd":7.40403,"snow_depth":0,"clouds":0,"ts":1607742060,"sunrise_ts":1607763972,"app_min_temp":17.3,"wind_spd":3.58314,"pop":0,"wind_cdir_full":"south-southeast","slp":1008.42,"moon_phase_lunation":0.97,"valid_date":"2020-12-12","app_max_temp":29.4,"vis":24.135,"dewpt":1.4,"snow":0,"uv":11.9145,"weather":{"icon":"c01d","code":800,"description":"Clear Sky"},"wind_dir":152,"max_dhi":null,"clouds_hi":0,"precip":0,"low_temp":18.3,"max_temp":31.8,"moonset_ts":1607812133,"datetime":"2020-12-12","temp":24.5,"min_temp":17.3,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1607849702,"wind_cdir":"SW","rh":28,"pres":952.284,"high_temp":32.6,"sunset_ts":1607901810,"ozone":271.946,"moon_phase":0.00141231,"wind_gust_spd":17.9229,"snow_depth":0,"clouds":0,"ts":1607828460,"sunrise_ts":1607850390,"app_min_temp":19.4,"wind_spd":4.91587,"pop":0,"wind_cdir_full":"southwest","slp":1009.08,"moon_phase_lunation":1,"valid_date":"2020-12-13","app_max_temp":26.9,"vis":24.1351,"dewpt":3.4,"snow":0,"uv":10.9545,"weather":{"icon":"c01d","code":800,"description":"Clear Sky"},"wind_dir":225,"max_dhi":null,"clouds_hi":0,"precip":0,"low_temp":17.6,"max_temp":28.6,"moonset_ts":1607902736,"datetime":"2020-12-13","temp":23.2,"min_temp":19.6,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1607939469,"wind_cdir":"E","rh":23,"pres":951.606,"high_temp":29.1,"sunset_ts":1607988248,"ozone":276.604,"moon_phase":0.0254232,"wind_gust_spd":10.5005,"snow_depth":0,"clouds":0,"ts":1607914860,"sunrise_ts":1607936809,"app_min_temp":17.6,"wind_spd":4.8597,"pop":0,"wind_cdir_full":"east","slp":1009.32,"moon_phase_lunation":0.03,"valid_date":"2020-12-14","app_max_temp":30.4,"vis":24.135,"dewpt":1.9,"snow":0,"uv":10.9188,"weather":{"icon":"c01d","code":800,"description":"Clear Sky"},"wind_dir":90,"max_dhi":null,"clouds_hi":0,"precip":0,"low_temp":20,"max_temp":33,"moonset_ts":1607993016,"datetime":"2020-12-14","temp":25,"min_temp":17.1,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1608029559,"wind_cdir":"S","rh":28,"pres":950.002,"high_temp":31.8,"sunset_ts":1608074685,"ozone":277.171,"moon_phase":0.0757318,"wind_gust_spd":18.0016,"snow_depth":0,"clouds":5,"ts":1608001260,"sunrise_ts":1608023230,"app_min_temp":19,"wind_spd":5.33089,"pop":0,"wind_cdir_full":"south","slp":1006.47,"moon_phase_lunation":0.07,"valid_date":"2020-12-15","app_max_temp":27.5,"vis":24.135,"dewpt":4,"snow":0,"uv":10.5382,"weather":{"icon":"c02d","code":801,"description":"Few clouds"},"wind_dir":173,"max_dhi":null,"clouds_hi":5,"precip":0,"low_temp":17.1,"max_temp":29.4,"moonset_ts":1608082822,"datetime":"2020-12-15","temp":24.2,"min_temp":19.1,"clouds_mid":0,"clouds_low":0},{"moonrise_ts":1608119810,"wind_cdir":"S","rh":24,"pres":952.9,"high_temp":29.4,"sunset_ts":1608161120,"ozone":279.053,"moon_phase":0.146602,"wind_gust_spd":4.70723,"snow_depth":0,"clouds":2,"ts":1608087660,"sunrise_ts":1608109652,"app_min_temp":17.1,"wind_spd":2.47178,"pop":20,"wind_cdir_full":"south","slp":1009.58,"moon_phase_lunation":0.1,"valid_date":"2020-12-16","app_max_temp":29.7,"vis":24.1351,"dewpt":2.1,"snow":0,"uv":10.896,"weather":{"icon":"c02d","code":801,"description":"Few clouds"},"wind_dir":174,"max_dhi":null,"clouds_hi":0,"precip":0.0625,"low_temp":18.3,"max_temp":32.3,"moonset_ts":1608172119,"datetime":"2020-12-16","temp":24.4,"min_temp":17,"clouds_mid":1,"clouds_low":1},{"moonrise_ts":1608210048,"wind_cdir":"S","rh":57,"pres":949.048,"high_temp":26,"sunset_ts":1608247554,"ozone":273.745,"moon_phase":0.231904,"wind_gust_spd":13.8178,"snow_depth":0,"clouds":39,"ts":1608174060,"sunrise_ts":1608196076,"app_min_temp":18.3,"wind_spd":4.15638,"pop":95,"wind_cdir_full":"south","slp":1006.5,"moon_phase_lunation":0.13,"valid_date":"2020-12-17","app_max_temp":29.4,"vis":17.5629,"dewpt":13.7,"snow":0,"uv":8.04565,"weather":{"icon":"t03d","code":202,"description":"Thunderstorm with heavy rain"},"wind_dir":178,"max_dhi":null,"clouds_hi":31,"precip":44.25,"low_temp":16.7,"max_temp":29.9,"moonset_ts":1608174564,"datetime":"2020-12-17","temp":23.4,"min_temp":17.5,"clouds_mid":32,"clouds_low":24}],"city_name":"Cordoba","lon":"-64.18105","timezone":"America/Argentina/Cordoba","lat":"-31.4135","country_code":"AR","state_code":"05"}`

func mockHttp(url string, method string, status int, jsonBody string) {
	mock := &rest.Mock{
		URL:          url,
		HTTPMethod:   method,
		ReqHeaders:   make(http.Header),
		RespHTTPCode: status,
		RespBody:     jsonBody,
	}

	_ = rest.AddMockups(mock)
}

func TestBeerForecast(t *testing.T) {
	handler := func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "localhost/v2.0/forecast/daily")
		rw.Write([]byte(weatherBitMockResponse))
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	os.Setenv("WEATHERBIT_URL", server.URL)
	defer os.Unsetenv("WEATHERBIT_URL")

	router := loadDependencies()

	request, err := http.NewRequest(http.MethodGet, "/beer-forecast?country=argentina&state=cordoba&city=cordoba&attendees=10&forecast_days=10&pack_units=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, request)
	fmt.Println(string(rr.Body.Bytes()))
	assert.Equal(t, 200, rr.Code)
}
