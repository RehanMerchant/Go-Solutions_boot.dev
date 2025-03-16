package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespLocation, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {

		locationRes := RespLocation{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return RespLocation{}, err
		}

		fmt.Println("Returing from cache")
		return locationRes, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocation{}, err
	}

	locationRes := RespLocation{}

	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, data)
	return locationRes, nil

}
