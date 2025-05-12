package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
	}

	data, ok := c.cache.Get(fullURL)
	if ok {
		// Cache hit, unmarshal the data and return it
		fmt.Println("Cache hit!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationAreasResp)

		if err != nil {
			return LocationAreasResp{}, err
		}

		return locationAreasResp, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("error: %s", resp.Status)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationAreasResp)

	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreasResp, nil
}


func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// Cache hit, unmarshal the data and return it
		fmt.Println("Cache hit!")
		locationArea := LocationArea{}
		err := json.Unmarshal(data, &locationArea)

		if err != nil {
			return LocationArea{}, err
		}

		return locationArea, nil
	}
	fmt.Println("Cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("error: %s", resp.Status)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(data, &locationArea)

	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationArea, nil
}
