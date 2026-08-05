package weathergo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type WeatherResponse struct {
	Latitude             float64     `json:"latitude"`
	Longitude            float64     `json:"longitude"`
	GenerationtimeMs     float64     `json:"generationtime_ms"`
	UtcOffsetSeconds     int         `json:"utc_offset_seconds"`
	Timezone             string      `json:"timezone"`
	TimezoneAbbreviation string      `json:"timezone_abbreviation"`
	Elevation            float64     `json:"elevation"`
	HourlyUnits          HourlyUnits `json:"hourly_units"`
	Hourly               HourlyData  `json:"hourly"`
	DailyUnits           DailyUnits  `json:"daily_units"`
	Daily                DailyData   `json:"daily"`
}

type HourlyUnits struct {
	Time               string `json:"time"`
	Temperature2m      string `json:"temperature_2m"`
	Precipitation      string `json:"precipitation"`
	WindSpeed10m       string `json:"wind_speed_10m"`
	WeatherCode        string `json:"weather_code"`
	RelativeHumidity2m string `json:"relative_humidity_2m"`
}

type HourlyData struct {
	Time               []string  `json:"time"`
	Temperature2m      []float64 `json:"temperature_2m"`
	Precipitation      []float64 `json:"precipitation"`
	WindSpeed10m       []float64 `json:"wind_speed_10m"`
	WeatherCode        []int     `json:"weather_code"`
	RelativeHumidity2m []int     `json:"relative_humidity_2m"`
}

type DailyUnits struct {
	Time               string `json:"time"`
	WeatherCode        string `json:"weather_code"`
	Temperature2mMax   string `json:"temperature_2m_max"`
	Temperature2mMin   string `json:"temperature_2m_min"`
	PrecipitationSum   string `json:"precipitation_sum"`
	PrecipitationHours string `json:"precipitation_hours"`
}

type DailyData struct {
	Time               []string  `json:"time"`
	WeatherCode        []int     `json:"weather_code"`
	Temperature2mMax   []float64 `json:"temperature_2m_max"`
	Temperature2mMin   []float64 `json:"temperature_2m_min"`
	PrecipitationSum   []float64 `json:"precipitation_sum"`
	PrecipitationHours []float64 `json:"precipitation_hours"`
}

// 先24時間分のデータを格納するための構造体
type HourlyForecast24h struct {
	Times              []string
	Temperatures       []float64
	Precipitations     []float64
	WindSpeeds         []float64
	WeatherCodes       []int
	RelativeHumidities []int
}

func GetWeatherInfo(latitude, longitude float64, useJma bool) (*WeatherResponse, error) {
	fallbackTimeZone := "UTC"
	if useJma {
		fallbackTimeZone = "Asia/Tokyo"
	}
	timeZoneParam := url.QueryEscape(getTimeZoneName(fallbackTimeZone))
	jmaParam := ""
	if useJma {
		jmaParam = "&models=jma_seamless"
	}
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%g&longitude=%g&daily=weather_code,temperature_2m_max,temperature_2m_min,precipitation_sum,precipitation_hours&hourly=temperature_2m,precipitation,wind_speed_10m,weather_code,relative_humidity_2m&timezone=%s%s", latitude, longitude, timeZoneParam, jmaParam)

	// 1. APIリクエストの送信
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

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

// HourlyDataから現在時刻以降の24時間分を抽出する関数
func GetNext24Hours(hourly HourlyData, timezone string) (*HourlyForecast24h, error) {
	// 現在時刻（ローカルタイム、またはAPIのタイムゾーンに合わせて調整）
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc = time.Local
	}
	now := time.Now().In(loc)

	startIndex := -1

	// time配列から「現在時刻以降（または直近の過去）」の開始インデックスを探す
	for i, tStr := range hourly.Time {
		// Open-Meteoの時刻フォーマット: "2026-02-04T00:00"
		t, err := time.ParseInLocation("2006-01-02T15:04", tStr, loc)
		if err != nil {
			continue
		}

		// 現在時刻以降、もしくはちょうど一致する時刻を最初の要素とする
		// 「現在時刻を含める」ために、時刻が未来または現在の最初のポイントを探す
		if !t.Before(now) {
			startIndex = i
			break
		}
	}

	// 見つからなかった場合や、残りが24時間未満の場合のフォールバック
	if startIndex == -1 {
		return &HourlyForecast24h{}, fmt.Errorf("No future data was found.")
	}

	// 24時間分（要素数24）を切り出し。データ全体の長さを超えないように調整
	endIndex := startIndex + 24
	if endIndex > len(hourly.Time) {
		endIndex = len(hourly.Time)
	}

	// スライスの切り出し
	return &HourlyForecast24h{
		Times:              hourly.Time[startIndex:endIndex],
		Temperatures:       hourly.Temperature2m[startIndex:endIndex],
		Precipitations:     hourly.Precipitation[startIndex:endIndex],
		WindSpeeds:         hourly.WindSpeed10m[startIndex:endIndex],
		WeatherCodes:       hourly.WeatherCode[startIndex:endIndex],
		RelativeHumidities: hourly.RelativeHumidity2m[startIndex:endIndex],
	}, nil
}

func getTimeZoneName(fallbackTimeZone string) string {
	// TZ環境変数が設定されていればそれを利用
	if tz := os.Getenv("TZ"); tz != "" {
		return tz
	}

	// Location名がIANAタイムゾーンなら利用
	if loc := time.Now().Location().String(); loc != "" && loc != "Local" {
		return loc
	}

	// フォールバック
	return fallbackTimeZone
}
