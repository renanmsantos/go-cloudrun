package usecases

type TemperaturesOutputDTO struct {
	CelsiusTemperature    float64 `json:"temp_C"`
	FahrenheitTemperature float64 `json:"temp_F"`
	KelvinTemperature     float64 `json:"temp_K"`
}

func NewTemperaturesOutputDTO(celsiusTemperature float64, fahrenheitTemperature float64, kelvinTemperature float64) TemperaturesOutputDTO {
	return TemperaturesOutputDTO{
		CelsiusTemperature:    celsiusTemperature,
		FahrenheitTemperature: fahrenheitTemperature,
		KelvinTemperature:     kelvinTemperature,
	}
}
