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
	Pokedex		map[string]PokemonResp
}

func NewClient() *Client {
	return &Client{
		BaseURL: 		"https://pokeapi.co/api/v2/",
		HttpClient: 	&http.Client{
			Timeout: 	time.Second * 10,
		},
		Cache:			pokecache.NewCache(time.Second * 5),
		Pokedex:		map[string]PokemonResp{},
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

func (c *Client) GetArea(areaName string) (AreaResp, error) {
	url := c.BaseURL + "location-area/" + areaName + "/"

	entry, exists := c.Cache.Get(url)
	if exists {
		return unmarshalArea(entry)
	}

	resp, err := c.HttpClient.Get(url)
	if err != nil {
		return AreaResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return AreaResp{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return AreaResp{}, err
	}
	c.Cache.Add(url, jsonData)
	return unmarshalArea(jsonData)
}

func unmarshalArea(jsonData []byte) (AreaResp, error) {
	var area AreaResp
	if err := json.Unmarshal(jsonData, &area); err != nil {
		return AreaResp{}, err
	}
	return area, nil
}

func (c *Client) GetPokemon(name string) (PokemonResp, error) {
	url := c.BaseURL + "pokemon/" + name + "/"

	entry, exists := c.Cache.Get(url)
	if exists {
		return unmarshalPokemon(entry)
	}

	resp, err := c.HttpClient.Get(url)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return PokemonResp{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}
	c.Cache.Add(url, jsonData)
	return unmarshalPokemon(jsonData)
}

func unmarshalPokemon(jsonData []byte) (PokemonResp, error) {
	var pkmn PokemonResp
	if err := json.Unmarshal(jsonData, &pkmn); err != nil {
		return PokemonResp{}, err
	}
	return pkmn, nil
}