package usecases

import (
	"errors"
	"log"
	"regexp"

	"github.com/renanmoreirasan/go-cloudrun/internal/gateways"
	"github.com/renanmoreirasan/go-cloudrun/internal/usecases/dtos"
)

func Execute(cep string) (dtos.TemperaturesOutput, error) {
	//CEP validation
	if !isValidCep(cep) {
		return dtos.TemperaturesOutput{}, errors.New("INVALID_CEP")
	}
	//Call API to get location
	location, err := gateways.GetLocation(cep)
	if err != nil {
		return dtos.TemperaturesOutput{}, err
	}
	//Call API to get temperature
	temperatures, err := gateways.GetLocationTemperature(location)
	if err != nil {
		return dtos.TemperaturesOutput{}, err
	}

	//Print temperature
	return dtos.NewTemperatures(
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
