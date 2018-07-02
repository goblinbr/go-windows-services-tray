package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Services []string
}

func ReadFile() (*Configuration, error) {
	file, err := os.Open("conf.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := &Configuration{}
	err = decoder.Decode(configuration)
	return configuration, err
}
