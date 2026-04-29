package api

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// APIレスポンスを格納する構造体
type Result struct {
	XMLName    xml.Name   `xml:"result"`
	Address    string     `xml:"address"`
	Coordinate coordinate `xml:"coordinate"`
	GoogleMaps string     `xml:"google_maps"`
}

type coordinate struct {
	Lat float64 `xml:"lat"`
	Lng float64 `xml:"lng"`
}

// 座標情報の取得
func GetCoordinate(location string) (*Result, error) {
	// パラメータチェック
	if location == "" {
		return nil, fmt.Errorf("argument error: location is empty")
	}

	// クエリパラメータをエンコード（日本語対応）
	apiURL := fmt.Sprintf("https://www.geocoding.jp/api/?q=%s", url.QueryEscape(location))

	// GETリクエストの送信
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

	// XMLのパース
	var result Result
	if err := xml.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
