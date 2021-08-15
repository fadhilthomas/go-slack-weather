package model

type CurrentWeatherResponse struct {
	Coord   `json:"coord"`
	Weather []Weather `json:"weather"`
	Main    `json:"main"`
	Wind    `json:"wind"`
	Rain    `json:"rain"`
	Clouds  `json:"clouds"`
	Dt      int    `json:"dt"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	Threehr int `json:"3h"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	Temp_min float64 `json:"temp_min"`
	Temp_max float64 `json:"temp_max"`
}
