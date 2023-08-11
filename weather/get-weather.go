package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type LocalNames struct {
	AR string `json:"ar"`
	DE string `json:"de"`
	KO string `json:"ko"`
	FR string `json:"fr"`
	RU string `json:"ru"`
	ES string `json:"es"`
	IT string `json:"it"`
	JA string `json:"ja"`
	EN string `json:"en"`
	UK string `json:"uk"`
	CR string `json:"cr"`
}

type Geo struct {
	Name       string     `json:"name"`
	LocalNames LocalNames `json:"local_names"`
	Lat        float64    `json:"lat"`
	Lon        float64    `json:"lon"`
	Country    string     `json:"country"`
	State      string     `json:"state"`
}

type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
	Gust  float32 `json:"gust"`
}

type Rain struct {
	OneH float32 `json:"1h"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Type    int    `json:"type"`
	Id      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

type OWResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Rain       Rain      `json:"rain"`
	Clouds     Clouds    `json:"clouds"`
	DT         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

func Getweather() {
	cityArg := os.Args[1:]
	var geo []Geo
	var weather OWResponse
	getGeo(cityArg, &geo)
	getWeather(geo, &weather)
	fmt.Println("+-------------+")
	fmt.Println("| Weather App |")
	fmt.Println("+-------------+")
	fmt.Println("Current temp in Edmonton: ", convKToC(weather.Main.Temp))
}

func getGeo(cityArg []string, geo *[]Geo) {
	geoUrl := "https://api.openweathermap.org/geo/1.0/direct?q=" + cityArg[0] + "&appid=" + os.Getenv("API_KEY")
	resp, err := http.Get(geoUrl)
	if err != nil {
		log.Printf("request Failed: %s", err)
		return
	}
	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &geo)
	if err != nil {
		log.Fatalln(err)
	}
}

func getWeather(geo []Geo, weather *OWResponse) {
	geoLat := strconv.FormatFloat(geo[0].Lat, 'f', -1, 64)
	geoLon := strconv.FormatFloat(geo[0].Lon, 'f', -1, 64)
	weatherUrl := "https://api.openweathermap.org/data/2.5/weather?lat=" + geoLat + "&lon=" + geoLon + "&appid=" + os.Getenv("API_KEY")
	resp, err := http.Get(weatherUrl)
	if err != nil {
		log.Printf("request Failed: %s", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &weather)
}

func convKToC(k float32) float32 {
	celsius := k - 273
	return celsius
}
