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
	CityInfo    cityInfo `json:"city_info"`
	CurrentInfo ctInfo   `json:"current_condition"`
	TodayWt     dayWt    `json:"fcst_day_0"`
	NextdayWt   dayWt    `json:"fcst_day_1"`
	AfterdayWt  dayWt    `json:"fcst_day_2"`
}

type ctInfo struct {
	Temp int `json:"tmp"`
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
	fmt.Println("START : Weather function - from weather command")
	var wt weatherInfo
	var wtInfo string
	var wtTdImg string
	var wtNxImg string
	var wtAfImg string
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
	wtAfImg = imgWt(wt.AfterdayWt.Condition)

	wtInfo = "Météo sur **" + wt.CityInfo.Name + "**:\n\nAujourd'hui :\nTempérature actuelle: " + strconv.Itoa(wt.CurrentInfo.Temp) + "\n" + wt.TodayWt.Condition + " " + wtTdImg + "\n" + "Min: " + strconv.Itoa(wt.TodayWt.Tmin) + "°\n" + "Max: " + strconv.Itoa(wt.TodayWt.Tmax) + "°\n\nPour demain :\n" + wt.NextdayWt.Condition + " " + wtNxImg + "\n" + "Min: " + strconv.Itoa(wt.NextdayWt.Tmin) + "°\n" + "Max: " + strconv.Itoa(wt.NextdayWt.Tmax) + "°\n\nEt pour " + wt.AfterdayWt.Day + "\n" + wt.AfterdayWt.Condition + " " + wtAfImg + "\n" + "Min: " + strconv.Itoa(wt.AfterdayWt.Tmin) + "°\n" + "Max: " + strconv.Itoa(wt.AfterdayWt.Tmax) + "°\n"

	return wtInfo, nil
}

func imgWt(cond string) (wtImg string) {

	switch cond {
	case "Ensoleillé":
		wtImg = "☀️"
	case "Eclaircies":
		wtImg = "🌤️"
	case "Ciel voilé", "Faibles passages nuageux", "Faiblement nuageux", "Développement nuageux":
		wtImg = "🌥️"
	case "Fortement nuageux":
		wtImg = "☁️"
	case "Brouillard", "Stratus":
		wtImg = "🌫️"
	case "Couvert avec averses":
		wtImg = "🌦️"
	case "Pluie faible", "Pluie modérée", "Pluie forte", "Averses de pluie faible", "Averses de pluie modérée", "Averses de pluie forte":
		wtImg = "🌧️"
	case "Faiblement orageux", "Orage modéré", "Fortement orageux":
		wtImg = "🌩️"
	case "Neige faible", "Neige modérée", "Neige forte":
		wtImg = "❄️"
	case "Pluie et neige mêlée faible", "Pluie et neige mêlée modérée", "Pluie et neige mêlée forte":
		wtImg = "🌨️"
	default:
		wtImg = ""

	}
	return wtImg
}
