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

func main() {
	var (
		pref = flag.Int("pref", 0, "Japanese prefecture code")
		days = flag.Int("days", 7, "N-Day Forecast, n must satisfy the condition 1 <= n <= 7")
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

	// データの表示
	fmt.Printf("Location: %s\n", coodinate.Place)
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
