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

type EffectArgs struct {
	Level    float64 `json:"level"`
	Duration float64 `json:"duration"`
	Ease     string  `json:"ease"`
}

type Effect struct {
	Time float64    `json:"time"`
	Type string     `json:"type"`
	Args EffectArgs `json:"args"`
}

type Chart struct {
	Notes   []Note   `json:"notes"`
	Effects []Effect `json:"effects"`
}

func UnmarshalChartFromFile(filename string) (*Chart, error) {
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
