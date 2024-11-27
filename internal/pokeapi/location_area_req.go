package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(locationUrl *string) (LocationAreasResp, error) {
	urlPath := "/location"
	fullUrl := baseUrl + urlPath

	if locationUrl != nil {
		fullUrl = *locationUrl
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	var location LocationAreasResp
	err = json.Unmarshal(body, &location)
	if err != nil {
		fmt.Println(err)
	}

	return location, nil
}
