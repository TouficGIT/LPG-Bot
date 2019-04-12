package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type weatherInfo struct {
	CityInfo  cityInfo `json:"city_info"`
	TodayWt   dayWt    `json:"fcst_day_0"`
	NextdayWt dayWt    `json:"fcst_day_1"`
}

type cityInfo struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type dayWt struct {
	Day       string `json:"day_long"`
	Tmin      int    `json:"tmin"`
	Tmax      int    `json:"tmax"`
	Condition string `json:"condition"`
}

// Weather func : retrieve the weather of a city.
// it takes the city and the day in input
// it return the weather of the city for the given day
func Weather(city string) (string, error) {

	var wt weatherInfo
	var wtInfo string
	var wtTdImg string
	var wtNxImg string
	city = strings.ToLower(city)

	resp, err := http.Get("https://www.prevision-meteo.ch/services/json/" + city)
	if err != nil {
		fmt.Println("Could not fetch weather infos")
		return "nil", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "nil", err
	}

	json.Unmarshal(body, &wt)
	wtTdImg = imgWt(wt.TodayWt.Condition)
	wtNxImg = imgWt(wt.NextdayWt.Condition)

	wtInfo = "Météo sur **" + wt.CityInfo.Name + "**:\n\nAujourd'hui :\n" + wt.TodayWt.Condition + " " + wtTdImg + "\n" + "Min: " + strconv.Itoa(wt.TodayWt.Tmin) + "°\n" + "Max: " + strconv.Itoa(wt.TodayWt.Tmax) + "°\n\nEt pour demain :\n" + wt.NextdayWt.Condition + " " + wtNxImg + "\n" + "Min: " + strconv.Itoa(wt.NextdayWt.Tmin) + "°\n" + "Max: " + strconv.Itoa(wt.NextdayWt.Tmax) + "°\n"

	return wtInfo, nil
}

func imgWt(cond string) (wtImg string) {

	switch cond {
	case "Ensoleillé":
		wtImg = "☀️"
	case "Eclaircies":
		wtImg = "🌤️"
	case "Ciel voilé":
		wtImg = "🌥️"
	case "Faiblement nuageux":
		wtImg = "☁️"
	case "Brouillard":
		wtImg = "🌫️"
	case "Pluie faible", "Pluie modérée", "Pluie forte", "Averses de pluie faible", "Averses de pluie modérée", "Averses de pluie forte":
		wtImg = "🌧️"
	case "Faiblement orageux", "Orage modéré", "Fortement orageux":
		wtImg = "🌩️"
	case "Neige faible", "Neige modérée", "Neige forte":
		wtImg = "❄️"
	default:
		wtImg = ""

	}
	return wtImg
}
