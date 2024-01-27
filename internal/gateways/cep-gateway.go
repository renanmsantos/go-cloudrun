package gateways

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Location struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

func GetLocation(cep string) (Location, error) {
	req, err := http.NewRequest("GET", "https://cep.awesomeapi.com.br/json/"+cep, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	if res.StatusCode == http.StatusNotFound {
		return Location{}, errors.New("CEP_NOT_FOUND")
	}
	defer res.Body.Close()

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var coordinates Location
	err = json.Unmarshal(resp, &coordinates)
	if err != nil {
		return Location{}, err
	}
	return coordinates, nil

}
