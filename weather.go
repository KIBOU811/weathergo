package weathergo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherResponse struct {
	Latitude             float64      `json:"latitude"`
	Longitude            float64      `json:"longitude"`
	GenerationTimeMS     float64      `json:"generationtime_ms"`
	UtcOffsetSeconds     int          `json:"utc_offset_seconds"`
	Timezone             string       `json:"timezone"`
	TimezoneAbbreviation string       `json:"timezone_abbreviation"`
	Elevation            float64      `json:"elevation"`
	CurrentUnits         CurrentUnits `json:"current_units"`
	Current              CurrentData  `json:"current"`
	DailyUnits           DailyUnits   `json:"daily_units"`
	Daily                DailyData    `json:"daily"`
}

type CurrentUnits struct {
	Time               string `json:"time"`
	Interval           string `json:"interval"`
	Temperature2m      string `json:"temperature_2m"`
	WeatherCode        string `json:"weather_code"`
	Precipitation      string `json:"precipitation"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
	PressureMsl        string `json:"pressure_msl"`
}

type CurrentData struct {
	Time               string  `json:"time"`
	Interval           int     `json:"interval"`
	Temperature2m      float64 `json:"temperature_2m"`
	WeatherCode        int     `json:"weather_code"`
	Precipitation      float64 `json:"precipitation"`
	RelativeHumidity2m int     `json:"relative_humidity_2m"`
	PressureMsl        float64 `json:"pressure_msl"`
}

type DailyUnits struct {
	Time                        string `json:"time"`
	WeatherCode                 string `json:"weather_code"`
	Temperature2mMax            string `json:"temperature_2m_max"`
	Temperature2mMin            string `json:"temperature_2m_min"`
	PrecipitationProbabilityMax string `json:"precipitation_probability_max"`
	PrecipitationSum            string `json:"precipitation_sum"`
	PressureMslMax              string `json:"pressure_msl_max"`
	PressureMslMin              string `json:"pressure_msl_min"`
	Sunrise                     string `json:"sunrise"`
	Sunset                      string `json:"sunset"`
}

type DailyData struct {
	Time                        []string  `json:"time"`
	WeatherCode                 []int     `json:"weather_code"`
	Temperature2mMax            []float64 `json:"temperature_2m_max"`
	Temperature2mMin            []float64 `json:"temperature_2m_min"`
	PrecipitationProbabilityMax []int     `json:"precipitation_probability_max"`
	PrecipitationSum            []float64 `json:"precipitation_sum"`
	PressureMslMax              []float64 `json:"pressure_msl_max"`
	PressureMslMin              []float64 `json:"pressure_msl_min"`
	Sunrise                     []string  `json:"sunrise"`
	Sunset                      []string  `json:"sunset"`
}

func GetWeatherInfo(latitude, longitude float64) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%g&longitude=%g&daily=weather_code,temperature_2m_max,temperature_2m_min,precipitation_probability_max,precipitation_sum,pressure_msl_max,pressure_msl_min,sunrise,sunset&current=temperature_2m,weather_code,precipitation,relative_humidity_2m,pressure_msl&timezone=Asia%sTokyo", latitude, longitude, "%2F")

	// 1. APIリクエストの送信
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected status code: %d\n", resp.StatusCode)
	}

	// 2. JSONのデコード
	var weather WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return nil, err
	}

	return &weather, err
}
