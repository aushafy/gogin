package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AutoGenerated struct {
	Count int `json:"count"`
	Data  []struct {
		AppTemp     float64  `json:"app_temp"`
		Aqi         int      `json:"aqi"`
		CityName    string   `json:"city_name"`
		Clouds      int      `json:"clouds"`
		CountryCode string   `json:"country_code"`
		Datetime    string   `json:"datetime"`
		Dewpt       float64  `json:"dewpt"`
		Dhi         float64  `json:"dhi"`
		Dni         float64  `json:"dni"`
		ElevAngle   float64  `json:"elev_angle"`
		Ghi         float64  `json:"ghi"`
		Gust        float64  `json:"gust"`
		HAngle      int      `json:"h_angle"`
		Lat         int      `json:"lat"`
		Lon         int      `json:"lon"`
		ObTime      string   `json:"ob_time"`
		Pod         string   `json:"pod"`
		Precip      int      `json:"precip"`
		Pres        float64  `json:"pres"`
		Rh          int      `json:"rh"`
		Slp         float64  `json:"slp"`
		Snow        int      `json:"snow"`
		SolarRad    float64  `json:"solar_rad"`
		Sources     []string `json:"sources"`
		StateCode   string   `json:"state_code"`
		Station     string   `json:"station"`
		Sunrise     string   `json:"sunrise"`
		Sunset      string   `json:"sunset"`
		Temp        int      `json:"temp"`
		Timezone    string   `json:"timezone"`
		Ts          int      `json:"ts"`
		Uv          float64  `json:"uv"`
		Vis         int      `json:"vis"`
		Weather     struct {
			Code        int    `json:"code"`
			Icon        string `json:"icon"`
			Description string `json:"description"`
		} `json:"weather"`
		WindCdir     string  `json:"wind_cdir"`
		WindCdirFull string  `json:"wind_cdir_full"`
		WindDir      int     `json:"wind_dir"`
		WindSpd      float64 `json:"wind_spd"`
	} `json:"data"`
}

func GetWeather(c *gin.Context) {
	url := "https://weatherbit-v1-mashape.p.rapidapi.com/current?lon=106&lat=-6"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", os.Getenv("API_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("API_HOST"))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	b := AutoGenerated{}
	err := json.Unmarshal(body, &b)
	if err != nil {
		fmt.Println("error unmarshal")
	}

	// fmt.Println(string(body))
	c.JSON(http.StatusOK, b)

}
