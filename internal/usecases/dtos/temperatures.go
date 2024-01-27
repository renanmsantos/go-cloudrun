package dtos

type TemperaturesOutput struct {
	CelsiusTemperature    float64 `json:"temp_C"`
	FahrenheitTemperature float64 `json:"temp_F"`
	KelvinTemperature     float64 `json:"temp_K"`
}

func NewTemperatures(celsiusTemperature float64, fahrenheitTemperature float64, kelvinTemperature float64) TemperaturesOutput {
	return TemperaturesOutput{CelsiusTemperature: celsiusTemperature, FahrenheitTemperature: fahrenheitTemperature, KelvinTemperature: kelvinTemperature}
}
