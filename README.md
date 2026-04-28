# weathergo

A module that retrieves weather information for a specified location using coordinates.

If the `pref` option is not specified, the application calculates coordinates from the IP address and retrieves weather information for those coordinates.

## Usage

```bash
$ go run ./cmd/console/main.go -days=3
Location: Asia/Tokyo
Current Weather: Cloud development not observed or not observable
Current Temp: 23.2°C, RH: 57%
Current Precip: 0.0mm, Pressure: 1008.2hPa

--- 3-Day Forecast ---
2026-04-28: Clouds generally forming or developing Max 23.2°C, Min 10.1°C, Precip Prob 0%
2026-04-29: Drizzle, not freezing, continuous, heavy Max 20.7°C, Min 14.4°C, Precip Prob 92%
2026-04-30: Drizzle, not freezing, continuous, moderate Max 16.9°C, Min 12.5°C, Precip Prob 67%
$ go run ./cmd/console/main.go -pref=1 -days=5
Location: 北海道 札幌市
Current Weather: Clouds generally forming or developing
Current Temp: 7.3°C, RH: 69%
Current Precip: 0.0mm, Pressure: 1011.2hPa

--- 5-Day Forecast ---
2026-04-28: Drizzle, not freezing, continuous, slight Max 9.0°C, Min 4.4°C, Precip Prob 50%
2026-04-29: Clouds generally forming or developing Max 12.9°C, Min 2.6°C, Precip Prob 90%
2026-04-30: Clouds generally forming or developing Max 17.9°C, Min 3.3°C, Precip Prob 18%
2026-05-01: Drizzle, not freezing, continuous, moderate Max 17.2°C, Min 6.1°C, Precip Prob 43%
2026-05-02: Drizzle, not freezing, continuous, moderate Max 14.6°C, Min 7.5°C, Precip Prob 59%
```

## `pref` Option Code

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
