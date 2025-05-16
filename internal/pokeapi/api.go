package pokeapi

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"io"
	"bootdev/pokedex/internal/pokecache"
)

type Client struct {
	BaseURL		string
	HttpClient	*http.Client
	Cache		*pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		BaseURL: 		"https://pokeapi.co/api/v2/",
		HttpClient: 	&http.Client{
			Timeout: 	time.Second * 10,
		},
		Cache:			pokecache.NewCache(time.Second * 5),
	}
}

func (c *Client) GetLocationAreas(url string) (LocationAreasResp, error) {
	if url == "" {
		url = c.BaseURL + "location-area/"
	}

	entry, exists := c.Cache.Get(url)
	if exists {
		return unmarshalLocationAreas(entry)
	}

	resp, err := c.HttpClient.Get(url)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
        return LocationAreasResp{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
    }

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}
	c.Cache.Add(url, jsonData)
	return unmarshalLocationAreas(jsonData)
}

func unmarshalLocationAreas(jsonData []byte) (LocationAreasResp, error) {
	var locationAreas LocationAreasResp
	if err := json.Unmarshal(jsonData, &locationAreas); err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreas, nil
}