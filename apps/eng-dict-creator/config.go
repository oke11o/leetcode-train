package main

import "os"

type Config struct {
	StateFile string
	WorkDir   string
}

func initConfig() *Config {
	cfg := Config{}
	cfg.WorkDir = os.Getenv("WORK_PATH")
	cfg.StateFile = os.Getenv("WORK_FILE")

	return &cfg
}
