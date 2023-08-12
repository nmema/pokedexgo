package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (locationAreas, error) {
	endpont := "/location-area"
	fullURL := baseURL + endpont

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationAreas{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreas{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreas{}, err
	}

	location := locationAreas{}
	err = json.Unmarshal(data, &location)
	if err != nil {
		return locationAreas{}, err
	}

	return location, nil
}
