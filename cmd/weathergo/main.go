package main

import (
	"flag"
	"fmt"

	"github.com/KIBOU811/weathergo"
)

// 範囲確認
func checkDays(days *int) *int {
	if *days < 1 || *days > 7 {
		*days = 7
	}

	return days
}

// Show Weather Info
func showWeatherInfo(weather *weathergo.WeatherResponse, hourly *weathergo.HourlyForecast24h, place string, days *int, show24 *bool) {
	fmt.Printf("Location: %s\n", place)
	fmt.Printf("Current Weather: %s\n", weathergo.GetWeatherName(weather.Hourly.WeatherCode[0]))
	fmt.Printf("Current Temperature: %.1f%s, Relative Humidity: %d%s\n", weather.Hourly.Temperature2m[0], weather.HourlyUnits.Temperature2m, weather.Hourly.RelativeHumidity2m[0], weather.HourlyUnits.RelativeHumidity2m)
	fmt.Printf("Current Precipitation: %.1f%s, Wind Speeds: %.1f%s\n", weather.Hourly.Precipitation[0], weather.HourlyUnits.Precipitation, weather.Hourly.WindSpeed10m[0], weather.HourlyUnits.WindSpeed10m)

	if *show24 {
		fmt.Println("\n--- 24-Hour Forecast ---")
		for i, time := range hourly.Times[:24] {
			fmt.Printf("%s: %s Temperature %.1f%s, Relative Humidity %d%s, Precipitation %.1f%s, Wind Speeds %.1f%s\n",
				time,
				weathergo.GetWeatherName(hourly.WeatherCodes[i]),
				hourly.Temperatures[i],
				weather.HourlyUnits.Temperature2m,
				hourly.RelativeHumidities[i],
				weather.HourlyUnits.RelativeHumidity2m,
				hourly.Precipitations[i],
				weather.HourlyUnits.Precipitation,
				hourly.WindSpeeds[i],
				weather.HourlyUnits.WindSpeed10m,
			)
		}
	}

	fmt.Printf("\n--- %d-Day Forecast ---\n", *days)
	for i, date := range hourly.Times[:*days] {
		fmt.Printf("%s: %s Max %.1f%s, Min %.1f%s, Precip Sum %.1f%s, Precip Hours %.0f%s\n",
			date,
			weathergo.GetWeatherName(weather.Daily.WeatherCode[i]),
			weather.Daily.Temperature2mMax[i],
			weather.DailyUnits.Temperature2mMax,
			weather.Daily.Temperature2mMin[i],
			weather.DailyUnits.Temperature2mMin,
			weather.Daily.PrecipitationSum[i],
			weather.DailyUnits.PrecipitationSum,
			weather.Daily.PrecipitationHours[i],
			weather.DailyUnits.PrecipitationHours,
		)
	}
}

// 天候情報の表示（日本語）
func showWeatherInfoJa(weather *weathergo.WeatherResponse, hourly *weathergo.HourlyForecast24h, place string, days *int, show24 *bool) {
	fmt.Printf("場所: %s\n", place)
	fmt.Printf("現在の天気: %s\n", weathergo.GetWeatherNameJa(weather.Hourly.WeatherCode[0]))
	fmt.Printf("気温: %.1f%s, 湿度: %d%s\n", weather.Hourly.Temperature2m[0], weather.HourlyUnits.Temperature2m, weather.Hourly.RelativeHumidity2m[0], weather.HourlyUnits.RelativeHumidity2m)
	fmt.Printf("降水量: %.1f%s, 風速: %.1f%s\n", weather.Hourly.Precipitation[0], weather.HourlyUnits.Precipitation, weather.Hourly.WindSpeed10m[0], weather.HourlyUnits.WindSpeed10m)

	if *show24 {
		fmt.Println("\n--- 24時間の天気 ---")
		for i, time := range hourly.Times[:24] {
			fmt.Printf("%s: %s 気温 %.1f%s, 湿度 %d%s, 降水量 %.1f%s, 風速 %.1f%s\n",
				time,
				weathergo.GetWeatherNameJa(hourly.WeatherCodes[i]),
				hourly.Temperatures[i],
				weather.HourlyUnits.Temperature2m,
				hourly.RelativeHumidities[i],
				weather.HourlyUnits.RelativeHumidity2m,
				hourly.Precipitations[i],
				weather.HourlyUnits.Precipitation,
				hourly.WindSpeeds[i],
				weather.HourlyUnits.WindSpeed10m,
			)
		}
	}

	fmt.Printf("\n--- %d日間の天気 ---\n", *days)
	for i, date := range weather.Daily.Time[:*days] {
		fmt.Printf("%s: %s 最高気温 %.1f%s, 最低気温 %.1f%s, 総降水量 %.1f%s, 降水時間 %.0f%s\n",
			date,
			weathergo.GetWeatherNameJa(weather.Daily.WeatherCode[i]),
			weather.Daily.Temperature2mMax[i],
			weather.DailyUnits.Temperature2mMax,
			weather.Daily.Temperature2mMin[i],
			weather.DailyUnits.Temperature2mMin,
			weather.Daily.PrecipitationSum[i],
			weather.DailyUnits.PrecipitationSum,
			weather.Daily.PrecipitationHours[i],
			weather.DailyUnits.PrecipitationHours,
		)
	}
}

func main() {
	var (
		loc    = flag.String("l", "", "Keywords for obtaining weather information.")
		days   = flag.Int("days", 7, "N-Day Forecast, n must satisfy the condition 1 <= n <= 7")
		ja     = flag.Bool("ja", false, "Output weather name in Japanese.")
		pref   = flag.Int("pref", 0, "Japanese prefecture code")
		jma    = flag.Bool("jma", false, "Use JMA for prediction model.")
		show24 = flag.Bool("show24", false, "Show 24-Hour forecast.")
	)
	flag.Parse()
	coodinate, err := weathergo.GetCoodinate(*loc, *pref)
	if err != nil {
		fmt.Printf("Error coodinate data: %v\n", err)
		return
	}
	weather, err := weathergo.GetWeatherInfo(coodinate.Latitude, coodinate.Longitude, *jma)
	if err != nil {
		fmt.Printf("Error weather data: %v\n", err)
		return
	}
	hourlyWeather24h, err := weathergo.GetNext24Hours(weather.Hourly, weather.Timezone)
	if err != nil {
		fmt.Printf("Error weather data: %v\n", err)
		return
	}
	days = checkDays(days)

	if *ja {
		showWeatherInfoJa(weather, hourlyWeather24h, coodinate.Place, days, show24)
	} else {
		showWeatherInfo(weather, hourlyWeather24h, coodinate.Place, days, show24)
	}
}
