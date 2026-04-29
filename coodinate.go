package weathergo

import (
	"fmt"
	"strconv"

	"github.com/KIBOU811/weathergo/api"
)

// 座標の定義
type Coodinate struct {
	Latitude  float64
	Longitude float64
	Place     string
}

// IPアドレスから座標の取得
func GetIpAddrCoodinate() (*Coodinate, error) {
	geo, err := api.GetGeoInfo()
	if err != nil {
		return nil, err
	}
	latitude, _ := strconv.ParseFloat(geo.Latitude, 64)
	longitude, _ := strconv.ParseFloat(geo.Longitude, 64)
	return &Coodinate{Latitude: latitude, Longitude: longitude, Place: geo.Timezone}, nil
}

// 住所などの地名から座標の取得
func GetAddressCoodinate(location string) (*Coodinate, error) {
	geo, err := api.GetCoordinate(location)
	if err != nil {
		return nil, err
	}
	return &Coodinate{Latitude: geo.Coordinate.Lat, Longitude: geo.Coordinate.Lng, Place: geo.GoogleMaps}, nil
}

// 座標の取得
func GetCoodinate(loc string, pref int) (*Coodinate, error) {
	// 地名から取得
	if loc != "" {
		coodinate, err := GetAddressCoodinate(loc)
		if err == nil {
			return coodinate, nil
		}
	}
	// IPアドレスから取得
	if pref > 47 || pref < 1 {
		return GetIpAddrCoodinate()
	}
	// 都道府県コードから取得
	location, _ := GetPrefCapitalData(pref)
	return &Coodinate{Latitude: location.Latitude, Longitude: location.Longitude, Place: fmt.Sprintf("%s %s", location.Prefecture, location.City)}, nil
}
