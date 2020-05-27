package main

import (
	"encoding/xml"
	"net/http"
	"text/template"

	"github.com/miuer/ncepu-work/architecture/com-experiment/soa/client"
)

// WeatherInfo -
type WeatherInfo struct {
	Information []string `xml:"string"`
}

// Weather -
type Weather struct {
	Province  string
	City      string
	QueryTime string
	Forecast  []Forecast
	Metar     string
	Tip       string
}

// Forecast -
type Forecast struct {
	Date        string
	Temperature string
	Wind        string
}

func main() {

	http.HandleFunc("/index", index)
	http.HandleFunc("/queryWeatherByLocation", handler)
	http.ListenAndServe(":10010", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./view/index.html"))

	err := t.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	location := r.PostFormValue("location")

	XMLweatherInfo := client.Run(location)

	weatherInfo := &WeatherInfo{}

	xml.Unmarshal([]byte(XMLweatherInfo), weatherInfo)

	weather := &Weather{
		Province:  weatherInfo.Information[0],
		City:      weatherInfo.Information[1],
		QueryTime: weatherInfo.Information[4],
		Metar:     weatherInfo.Information[10],
		Tip:       weatherInfo.Information[11],
	}

	f1 := Forecast{
		Date:        weatherInfo.Information[6],
		Temperature: weatherInfo.Information[5],
		Wind:        weatherInfo.Information[7],
	}

	f2 := Forecast{
		Date:        weatherInfo.Information[13],
		Temperature: weatherInfo.Information[12],
		Wind:        weatherInfo.Information[14],
	}

	f3 := Forecast{
		Date:        weatherInfo.Information[18],
		Temperature: weatherInfo.Information[17],
		Wind:        weatherInfo.Information[19],
	}

	weather.Forecast = append(weather.Forecast, f1, f2, f3)

	// w.Write([]byte(weather))

	//	log.Println(weather)

	t := template.Must(template.ParseFiles("./view/weather.html"))

	t.Execute(w, weather)

}
