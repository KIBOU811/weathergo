package weathergo

import (
	"fmt"
	"strconv"
)

// 座標の定義
type Coodinate struct {
	Latitude  float64
	Longitude float64
	Place     string
}

// IPアドレスから座標の取得
func GetIpAddrCoodinate() (*Coodinate, error) {
	geo, err := GetGeoInfo()
	if err != nil {
		return nil, err
	}
	latitude, _ := strconv.ParseFloat(geo.Latitude, 64)
	longitude, _ := strconv.ParseFloat(geo.Longitude, 64)
	return &Coodinate{Latitude: latitude, Longitude: longitude, Place: geo.Timezone}, nil
}

// 座標の取得
func GetCoodinate(pref int) (*Coodinate, error) {
	if pref > 47 || pref < 1 {
		return GetIpAddrCoodinate()
	}
	location, _ := GetPrefCapitalData(pref)
	return &Coodinate{Latitude: location.Latitude, Longitude: location.Longitude, Place: fmt.Sprintf("%s %s", location.Prefecture, location.City)}, nil
}
