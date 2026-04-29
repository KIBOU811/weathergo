package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// GeoJS APIからのレスポンス構造
type GeoResponse struct {
	IP           string `json:"ip"`
	Country      string `json:"country"`
	CountryCode  string `json:"country_code"`
	Region       string `json:"region"`
	City         string `json:"city"`
	Latitude     string `json:"latitude"`  // 文字列で返るためstring
	Longitude    string `json:"longitude"` // 文字列で返るためstring
	Organization string `json:"organization"`
	Timezone     string `json:"timezone"`
}

// 座標情報の取得
func GetGeoInfo() (*GeoResponse, error) {
	apiURL := "https://get.geojs.io/v1/ip/geo.json"

	// HTTP GETリクエストの送信
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスボディの読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// JSONのパース（デコード）
	var geo GeoResponse
	if err := json.Unmarshal(body, &geo); err != nil {
		return nil, err
	}
	if _, err := strconv.ParseFloat(geo.Latitude, 64); err != nil {
		return nil, err
	}
	if _, err := strconv.ParseFloat(geo.Longitude, 64); err != nil {
		return nil, err
	}

	return &geo, err
}
