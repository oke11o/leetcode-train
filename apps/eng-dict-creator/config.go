package main

import "os"

type Config struct {
	StateFile    string
	WorkDir      string
	TextFilepath string
	LingvoApiKey string
}

func initConfig() *Config {
	cfg := Config{}
	cfg.WorkDir = os.Getenv("WORK_PATH")
	cfg.StateFile = os.Getenv("WORK_FILE")
	cfg.TextFilepath = os.Getenv("TEXT_FILEPATH")
	cfg.LingvoApiKey = os.Getenv("LINGVO_API_KEY")

	return &cfg
}
