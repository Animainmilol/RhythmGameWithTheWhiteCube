package main

import (
	"encoding/json"
	"os"
)

type Note struct {
	Time     float64 `json:"time"`
	Movement string  `json:"movement"`
	Type     string  `json:"type"`
	Duration float64 `json:"duration"`
}

type ZoomEffect struct {
	Level    float64 `json:"level"`
	Duration float64 `json:"duration"`
	Ease     string  `json:"ease"`
}

type Effects struct {
	Zoom []ZoomEffect `json:"zoom"`
}

type Chart struct {
	Notes   []Note    `json:"notes"`
	Effects []Effects `json:"effects"`
}

func ParseChartFromFile(filename string) (*Chart, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var chart Chart
	err = json.Unmarshal(data, &chart)
	if err != nil {
		return nil, err
	}

	return &chart, nil
}
