package main

import "time"

type State struct {
	Name     string    `json:"name"`
	FileName string    `json:"file_name"`
	ModTime  time.Time `json:"mod_time"`
	Size     int64     `json:"size"`
}
