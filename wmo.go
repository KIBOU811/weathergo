package weathergo

// 気象通報式コード（00-99）を表す型
type WMOWeatherCode int

// A map linking WMOWeatherCode with weather text.
var weatherCodeMap = map[WMOWeatherCode]string{
	// 00-19: 降水、霧、砂嵐などがない状態
	0:  "Cloud development not observed or not observable",
	1:  "Clouds generally dissolving or becoming less developed",
	2:  "State of sky on the whole unchanged",
	3:  "Clouds generally forming or developing",
	4:  "Visibility reduced by smoke",
	5:  "Haze",
	6:  "Widespread dust in suspension in the air",
	7:  "Dust or sand raised by wind",
	8:  "Well developed dust whirl(s) or sand whirl(s)",
	9:  "Duststorm or sandstorm within sight",
	10: "Mist",
	11: "Patches shallow fog or ice fog",
	12: "More or less continuous shallow fog or ice fog",
	13: "Lightning visible, no thunder heard",
	14: "Precipitation within sight, not reaching the ground",
	15: "Precipitation within sight, reaching the ground, distant (> 5 km)",
	16: "Precipitation within sight, reaching the ground, near to but not at the station",
	17: "Thunderstorm, but no precipitation at the time of observation",
	18: "Squalls",
	19: "Funnel cloud(s) (Tornado or water-spout)",

	// 20-29: 前1時間以内に降水等があったが観測時はなし
	20: "Drizzle (not freezing) or snow grains (not as showers) in preceding hour",
	21: "Rain (not freezing) in preceding hour",
	22: "Snow in preceding hour",
	23: "Rain and snow or ice pellets in preceding hour",
	24: "Freezing drizzle or freezing rain in preceding hour",
	25: "Shower(s) of rain in preceding hour",
	26: "Shower(s) of snow, or of rain and snow in preceding hour",
	27: "Shower(s) of hail, or of rain and hail in preceding hour",
	28: "Fog or ice fog in preceding hour",
	29: "Thunderstorm (with or without precipitation) in preceding hour",

	// 30-39: 砂嵐、地吹雪
	30: "Slight or moderate duststorm or sandstorm - decreased during preceding hour",
	31: "Slight or moderate duststorm or sandstorm - no appreciable change",
	32: "Slight or moderate duststorm or sandstorm - begun or increased",
	33: "Severe duststorm or sandstorm - decreased during preceding hour",
	34: "Severe duststorm or sandstorm - no appreciable change",
	35: "Severe duststorm or sandstorm - begun or increased",
	36: "Slight or moderate blowing snow (low level)",
	37: "Heavy drifting snow",
	38: "Slight or moderate blowing snow (high level)",
	39: "Heavy drifting snow",

	// 40-49: 霧または氷霧
	40: "Fog or ice fog at a distance",
	41: "Fog or ice fog in patches",
	42: "Fog or ice fog, sky visible - thinner during preceding hour",
	43: "Fog or ice fog, sky invisible - thinner during preceding hour",
	44: "Fog or ice fog, sky visible - no change",
	45: "Fog or ice fog, sky invisible - no change",
	46: "Fog or ice fog, sky visible - begun or thicker",
	47: "Fog or ice fog, sky invisible - begun or thicker",
	48: "Fog, depositing rime, sky visible",
	49: "Fog, depositing rime, sky invisible",

	// 50-59: 霧雨 (Drizzle)
	50: "Drizzle, not freezing, intermittent, slight",
	51: "Drizzle, not freezing, continuous, slight",
	52: "Drizzle, not freezing, intermittent, moderate",
	53: "Drizzle, not freezing, continuous, moderate",
	54: "Drizzle, not freezing, intermittent, heavy",
	55: "Drizzle, not freezing, continuous, heavy",
	56: "Drizzle, freezing, slight",
	57: "Drizzle, freezing, moderate or heavy",
	58: "Drizzle and rain, slight",
	59: "Drizzle and rain, moderate or heavy",

	// 60-69: 雨 (Rain)
	60: "Rain, not freezing, intermittent, slight",
	61: "Rain, not freezing, continuous, slight",
	62: "Rain, not freezing, intermittent, moderate",
	63: "Rain, not freezing, continuous, moderate",
	64: "Rain, not freezing, intermittent, heavy",
	65: "Rain, not freezing, continuous, heavy",
	66: "Rain, freezing, slight",
	67: "Rain, freezing, moderate or heavy",
	68: "Rain or drizzle and snow, slight",
	69: "Rain or drizzle and snow, moderate or heavy",

	// 70-79: 雪・固形降水 (Solid precipitation)
	70: "Intermittent fall of snowflakes, slight",
	71: "Continuous fall of snowflakes, slight",
	72: "Intermittent fall of snowflakes, moderate",
	73: "Continuous fall of snowflakes, moderate",
	74: "Intermittent fall of snowflakes, heavy",
	75: "Continuous fall of snowflakes, heavy",
	76: "Diamond dust",
	77: "Snow grains",
	78: "Isolated star-like snow crystals",
	79: "Ice pellets",

	// 80-99: 驟雨（しゅうう）・雷雨 (Showers / Thunderstorms)
	80: "Rain shower(s), slight",
	81: "Rain shower(s), moderate or heavy",
	82: "Rain shower(s), violent",
	83: "Shower(s) of rain and snow mixed, slight",
	84: "Shower(s) of rain and snow mixed, moderate or heavy",
	85: "Snow shower(s), slight",
	86: "Snow shower(s), moderate or heavy",
	87: "Shower(s) of snow pellets or small hail, slight",
	88: "Shower(s) of snow pellets or small hail, moderate or heavy",
	89: "Shower(s) of hail, slight",
	90: "Shower(s) of hail, moderate or heavy",
	91: "Slight rain at time of observation, thunderstorm during preceding hour",
	92: "Moderate or heavy rain, thunderstorm during preceding hour",
	93: "Slight snow, or rain and snow mixed or hail, thunderstorm during preceding hour",
	94: "Moderate or heavy snow, or rain and snow mixed or hail, thunderstorm during preceding hour",
	95: "Thunderstorm, slight or moderate, without hail but with rain/snow",
	96: "Thunderstorm, slight or moderate, with hail",
	97: "Thunderstorm, heavy, without hail but with rain/snow",
	98: "Thunderstorm combined with duststorm or sandstorm",
	99: "Thunderstorm, heavy, with hail",
}

// WMOWeatherCode と天気のテキストを紐付けるマップ（日本語）
var weatherCodeMapJa = map[WMOWeatherCode]string{
	// 00–19：降水、霧、砂塵嵐などがない場合
	0:  "雲の発達が観測されない、または観測不能",
	1:  "雲が消えつつある、または衰弱しつつある",
	2:  "空の状態に変化がない",
	3:  "雲が形成されつつある、または発達しつつある",
	4:  "煙により視程が減少",
	5:  "煙霧（ヘイズ）",
	6:  "空中に浮遊した広範囲のじんあい",
	7:  "風で巻き上げられた塵や砂",
	8:  "発達した塵旋風または砂旋風",
	9:  "視界内の砂塵嵐、または前1時間以内の砂塵嵐",
	10: "もや（ミスト）",
	11: "低い霧または氷霧（パッチ状）",
	12: "低い霧または氷霧（連続的）",
	13: "雷光が見えるが、雷鳴は聞こえない",
	14: "視界内に降水があるが、地面に達していない",
	15: "遠方に降水が見える（5km以上）",
	16: "近傍に降水が見える（観測所にはかかっていない）",
	17: "雷雨（観測時に降水なし）",
	18: "突風（スクール）",
	19: "漏斗雲（竜巻または水竜巻）",

	// 20–29：前1時間以内に降水などがあったが、観測時にはない
	20: "前1時間に霧雨または霧雪",
	21: "前1時間に雨",
	22: "前1時間に雪",
	23: "前1時間にみぞれまたはアイスペレット",
	24: "前1時間に着氷性の霧雨または雨",
	25: "前1時間にしゅう雨（にわか雨）",
	26: "前1時間にしゅう雪、またはしゅう雨を伴う雪",
	27: "前1時間に雹（ひょう）",
	28: "前1時間に霧または氷霧",
	29: "前1時間に雷雨",

	// 30-39: 砂塵嵐、または地ふぶき
	30: "弱または並の砂塵嵐（弱まった）",
	31: "弱または並の砂塵嵐（変化なし）",
	32: "弱または並の砂塵嵐（強まった）",
	33: "強烈な砂塵嵐（弱まった）",
	34: "強烈な砂塵嵐（変化なし）",
	35: "強烈な砂塵嵐（強まった）",
	36: "弱または並の低い地ふぶき",
	37: "強い低い地ふぶき",
	38: "弱または並の高い地ふぶき",
	39: "強い高い地ふぶき",

	// 40-49: 霧または氷霧
	40: "遠方の霧",
	41: "霧または氷霧（パッチ状）",
	42: "霧（空が見える、薄くなった）",
	43: "霧（空が見えない、薄くなった）",
	44: "霧（空が見える、変化なし）",
	45: "霧（空が見えない、変化なし）",
	46: "霧（空が見える、濃くなった）",
	47: "霧（空が見えない、濃くなった）",
	48: "霧（樹氷を伴う、空が見える）",
	49: "霧（樹氷を伴う、空が見えない）",

	// 50-59: 霧雨
	50: "断続的な霧雨（弱い）",
	51: "連続的な霧雨（弱い）",
	52: "断続的な霧雨（並）",
	53: "連続的な霧雨（並）",
	54: "断続的な霧雨（強い）",
	55: "連続的な霧雨（強い）",
	56: "着氷性の霧雨（弱い）",
	57: "着氷性の霧雨（並または強い）",
	58: "霧雨を伴う雨（弱い）",
	59: "霧雨を伴う雨（並または強い）",

	// 60-69: 雨
	60: "断続的な雨（弱い）",
	61: "連続的な雨（弱い）",
	62: "断続的な雨（並）",
	63: "連続的な雨（並）",
	64: "断続的な雨（強い）",
	65: "連続的な雨（強い）",
	66: "着氷性の雨（弱い）",
	67: "着氷性の雨（並または強い）",
	68: "雨または霧雨を伴う雪（弱い）",
	69: "雨または霧雨を伴う雪（並または強い）",

	// 70-79: 固形降水
	70: "断続的な雪（弱い）",
	71: "連続的な雪（弱い）",
	72: "断続的な雪（並）",
	73: "連続的な雪（並）",
	74: "断続的な雪（強い）",
	75: "連続的な雪（強い）",
	76: "氷晶（ダイヤモンドダスト）",
	77: "霧雪",
	78: "孤立した星型の雪の結晶",
	79: "アイスペレット",

	// 80-99: しゅう雨性、または雷雨
	80: "しゅう雨（弱い）",
	81: "しゅう雨（並または強い）",
	82: "しゅう雨（激しい）",
	83: "みぞれのしゅう雨（弱い）",
	84: "みぞれのしゅう雨（並または強い）",
	85: "しゅう雪（弱い）",
	86: "しゅう雪（並または強い）",
	87: "雪あられまたは小さな雹（弱い）",
	88: "雪あられまたは小さな雹（並または強い）",
	89: "雹（弱い）",
	90: "雹（並または強い）",
	91: "弱い雨（前1時間以内に雷雨）",
	92: "並または強い雨（前1時間以内に雷雨）",
	93: "弱い雪・みぞれ・雹（前1時間以内に雷雨）",
	94: "並または強い雪・みぞれ・雹（前1時間以内に雷雨）",
	95: "雷雨（弱いまたは並、雨・雪を伴う）",
	96: "雷雨（弱いまたは並、雹を伴う）",
	97: "雷雨（強い、雨・雪を伴う）",
	98: "雷雨（砂塵嵐を伴う）",
	99: "雷雨（強い、雹を伴う）",
}

// Convert WMO codes to their corresponding strings.
func GetWeatherName(code int) string {
	if name, ok := weatherCodeMap[WMOWeatherCode(code)]; ok {
		return name
	}
	return "Unknown code"
}

// WMOコードを対応する日本語の文字列に変換する
func GetWeatherNameJa(code int) string {
	if name, ok := weatherCodeMapJa[WMOWeatherCode(code)]; ok {
		return name
	}
	return "不明なコード"
}
