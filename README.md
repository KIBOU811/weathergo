# weathergo

A module that retrieves weather information for a specified location using coordinates.

The `l` option retrieves the coordinates of a specified location.  
If neither the `l` option nor the `pref` option is specified, the application calculates coordinates from the IP address and retrieves weather information for those coordinates.

## Usage

```bash
$ go run ./cmd/weathergo/main.go -l 札幌 -days 3
Location: 北海道札幌市
Current Weather: Clouds generally forming or developing
Current Temp: 13.2°C, RH: 44%
Current Precip: 0.0mm, Pressure: 1012.2hPa

--- 3-Day Forecast ---
2026-04-29: Rain, not freezing, continuous, slight Max 13.2°C, Min 3.6°C, Precip Prob 94%
2026-04-30: Drizzle, not freezing, continuous, slight Max 18.1°C, Min 3.3°C, Precip Prob 66%
2026-05-01: Rain, not freezing, continuous, slight Max 15.1°C, Min 6.6°C, Precip Prob 100%
$ go run ./cmd/weathergo/main.go -pref=1 -days=5 -ja
場所: 北海道 札幌市
現在の天気: 空の状態に変化がない
気温: 4.5°C, 湿度: 81%
降水量: 0.0mm, 気圧: 1013.6hPa

--- 5日間の天気 ---
2026-04-28: 連続的な霧雨（弱い） 最高気温 9.0°C, 最低気温 4.3°C, 降水確率 95%
2026-04-29: 雲が形成されつつある、または発達しつつある 最高気温 12.9°C, 最低気温 2.9°C, 降水確率 96%
2026-04-30: 連続的な霧雨（弱い） 最高気温 17.8°C, 最低気温 2.9°C, 降水確率 29%
2026-05-01: 連続的な雨（弱い） 最高気温 16.9°C, 最低気温 5.9°C, 降水確率 60%
2026-05-02: 連続的な雨（弱い） 最高気温 13.1°C, 最低気温 7.1°C, 降水確率 78%
```

### `pref` Option Code

| コード | 都道府県 | 都道府県庁所在地 |
| :--- | :--- | :--- |
| 1 | 北海道 | 札幌市 |
| 2 | 青森県 | 青森市 |
| 3 | 岩手県 | 盛岡市 |
| 4 | 宮城県 | 仙台市 |
| 5 | 秋田県 | 秋田市 |
| 6 | 山形県 | 山形市 |
| 7 | 福島県 | 福島市 |
| 8 | 茨城県 | 水戸市 |
| 9 | 栃木県 | 宇都宮市 |
| 10 | 群馬県 | 前橋市 |
| 11 | 埼玉県 | さいたま市 |
| 12 | 千葉県 | 千葉市 |
| 13 | 東京都 | 新宿区 |
| 14 | 神奈川県 | 横浜市 |
| 15 | 新潟県 | 新潟市 |
| 16 | 富山県 | 富山市 |
| 17 | 石川県 | 金沢市 |
| 18 | 福井県 | 福井市 |
| 19 | 山梨県 | 甲府市 |
| 20 | 長野県 | 長野市 |
| 21 | 岐阜県 | 岐阜市 |
| 22 | 静岡県 | 静岡市 |
| 23 | 愛知県 | 名古屋市 |
| 24 | 三重県 | 津市 |
| 25 | 滋賀県 | 大津市 |
| 26 | 京都府 | 京都市 |
| 27 | 大阪府 | 大阪市 |
| 28 | 兵庫県 | 神戸市 |
| 29 | 奈良県 | 奈良市 |
| 30 | 和歌山県 | 和歌山市 |
| 31 | 鳥取県 | 鳥取市 |
| 32 | 島根県 | 松江市 |
| 33 | 岡山県 | 岡山市 |
| 34 | 広島県 | 広島市 |
| 35 | 山口県 | 山口市 |
| 36 | 徳島県 | 徳島市 |
| 37 | 香川県 | 高松市 |
| 38 | 愛媛県 | 松山市 |
| 39 | 高知県 | 高知市 |
| 40 | 福岡県 | 福岡市 |
| 41 | 佐賀県 | 佐賀市 |
| 42 | 長崎県 | 長崎市 |
| 43 | 熊本県 | 熊本市 |
| 44 | 大分県 | 大分市 |
| 45 | 宮崎県 | 宮崎市 |
| 46 | 鹿児島県 | 鹿児島市 |
| 47 | 沖縄県 | 那覇市 |

## Data Sources & APIs

This project retrieves weather and location data using the following services:

* **[GeoJS](https://www.geojs.io/)** - IP-based Geolocation API for automatic location detection.
* **[Geocoding.jp API](https://geocoding.jp/api/)** - For converting Japanese address queries into geographic coordinates.
* **[Open-Meteo](https://open-meteo.com/)** - Weather forecast data.
  * *Note: Weather data is provided by Open-Meteo.com under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/).*
