package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	FromFile  string `yaml:"fromFile"`
	ToFile    string `yaml:"toFile"`
	Increment int64  `yaml:"increment"` // miliseconds
}

func readConfig() (Config, error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return Config{}, fmt.Errorf("read file err: %v", err)
	}
	c := Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return Config{}, fmt.Errorf("marshal err: %v", err)
	}

	return c, nil
}
