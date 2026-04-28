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
func showWeatherInfo(weather *weathergo.WeatherResponse, place string, days *int) {
	fmt.Printf("Location: %s\n", place)
	fmt.Printf("Current Weather: %s\n", weathergo.GetWeatherName(weather.Current.WeatherCode))
	fmt.Printf("Current Temp: %.1f%s, RH: %d%s\n", weather.Current.Temperature2m, weather.CurrentUnits.Temperature2m, weather.Current.RelativeHumidity2m, weather.CurrentUnits.RelativeHumidity2m)
	fmt.Printf("Current Precip: %.1f%s, Pressure: %.1f%s\n", weather.Current.Precipitation, weather.CurrentUnits.Precipitation, weather.Current.PressureMsl, weather.CurrentUnits.PressureMsl)

	fmt.Printf("\n--- %d-Day Forecast ---\n", *days)
	for i, date := range weather.Daily.Time[:*days] {
		fmt.Printf("%s: %s Max %.1f%s, Min %.1f%s, Precip Prob %d%s\n",
			date,
			weathergo.GetWeatherName(weather.Daily.WeatherCode[i]),
			weather.Daily.Temperature2mMax[i],
			weather.DailyUnits.Temperature2mMax,
			weather.Daily.Temperature2mMin[i],
			weather.DailyUnits.Temperature2mMin,
			weather.Daily.PrecipitationProbabilityMax[i],
			weather.DailyUnits.PrecipitationProbabilityMax,
		)
	}
}

// 天候情報の表示（日本語）
func showWeatherInfoJa(weather *weathergo.WeatherResponse, place string, days *int) {
	fmt.Printf("場所: %s\n", place)
	fmt.Printf("現在の天気: %s\n", weathergo.GetWeatherNameJa(weather.Current.WeatherCode))
	fmt.Printf("気温: %.1f%s, 湿度: %d%s\n", weather.Current.Temperature2m, weather.CurrentUnits.Temperature2m, weather.Current.RelativeHumidity2m, weather.CurrentUnits.RelativeHumidity2m)
	fmt.Printf("降水量: %.1f%s, 気圧: %.1f%s\n", weather.Current.Precipitation, weather.CurrentUnits.Precipitation, weather.Current.PressureMsl, weather.CurrentUnits.PressureMsl)

	fmt.Printf("\n--- %d日間の天気 ---\n", *days)
	for i, date := range weather.Daily.Time[:*days] {
		fmt.Printf("%s: %s 最高気温 %.1f%s, 最低気温 %.1f%s, 降水確率 %d%s\n",
			date,
			weathergo.GetWeatherNameJa(weather.Daily.WeatherCode[i]),
			weather.Daily.Temperature2mMax[i],
			weather.DailyUnits.Temperature2mMax,
			weather.Daily.Temperature2mMin[i],
			weather.DailyUnits.Temperature2mMin,
			weather.Daily.PrecipitationProbabilityMax[i],
			weather.DailyUnits.PrecipitationProbabilityMax,
		)
	}
}

func main() {
	var (
		pref = flag.Int("pref", 0, "Japanese prefecture code")
		days = flag.Int("days", 7, "N-Day Forecast, n must satisfy the condition 1 <= n <= 7")
		ja   = flag.Bool("ja", false, "Output weather name in Japanese.")
	)
	flag.Parse()
	coodinate, err := weathergo.GetCoodinate(*pref)
	if err != nil {
		fmt.Printf("Error coodinate data: %v\n", err)
		return
	}
	weather, err := weathergo.GetWeatherInfo(coodinate.Latitude, coodinate.Longitude)
	if err != nil {
		fmt.Printf("Error weather data: %v\n", err)
		return
	}
	days = checkDays(days)

	if *ja {
		showWeatherInfoJa(weather, coodinate.Place, days)
	} else {
		showWeatherInfo(weather, coodinate.Place, days)
	}
}
