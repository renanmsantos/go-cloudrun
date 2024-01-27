package usecases

import (
	"errors"
	"log"
	"regexp"

	"github.com/renanmoreirasan/go-cloudrun/internal/gateways"
)

func Execute(cep string) (TemperaturesOutputDTO, error) {
	//CEP validation
	if !isValidCep(cep) {
		return TemperaturesOutputDTO{}, errors.New("INVALID_CEP")
	}
	//Call API to get location
	location, err := gateways.GetLocation(cep)
	if err != nil {
		return TemperaturesOutputDTO{}, err
	}
	//Call API to get temperature
	temperatures, err := gateways.GetLocationTemperature(location)
	if err != nil {
		return TemperaturesOutputDTO{}, err
	}

	//Print temperature
	return NewTemperaturesOutputDTO(
		temperatures.Current.Celsius,
		temperatures.Current.Fahrenheit,
		convertCelsiusToKelvin(temperatures.Current.Celsius),
	), nil
}

func isValidCep(cep string) bool {
	containsEightDigitsRegex := `^[0-9]{8}$`
	match, err := regexp.MatchString(containsEightDigitsRegex, cep)
	if err != nil {
		log.Printf("Error while validating CEP: %v", err)
		return false
	}
	return match
}

func convertCelsiusToKelvin(celsiusTemperature float64) float64 {
	return celsiusTemperature + 273.15
}
