package main

import (
	"encoding/json"
	"os"
)

type TilemapLayer struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type Tilemap struct {
	Layers []TilemapLayer `json:"layers"`
}

func NewTilemap(filepath string) (*Tilemap, error) {
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tilemap Tilemap
	err = json.Unmarshal(contents, &tilemap)
	if err != nil {
		return nil, err
	}

	return &tilemap, nil
}
