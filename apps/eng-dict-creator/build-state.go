package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func buildState(cfg *Config) error {
	entries, err := ioutil.ReadDir(cfg.WorkDir)
	if err != nil {
		return fmt.Errorf("cant read dir `%s`, err: %w", cfg.WorkDir, err)
	}

	result := map[string]State{}
	for _, fname := range entries {
		ext := extByName(fname.Name())
		if ext != "md" {
			continue
		}
		fpath := cfg.WorkDir + "/" + fname.Name()
		file, err := os.Open(fpath)
		if err != nil {
			return fmt.Errorf("cant read file `%s`, err: %w", fpath, err)
		}
		info, err := file.Stat()
		if err != nil {
			return fmt.Errorf("something went wrong for fileinfo `%s`, err: %w", fpath, err)
		}
		state := fileinfo2State(info)
		result[state.Name] = state
	}

	// save to file
	jsonStr, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshaling state map, err: %w", err)
	}

	outFile := cfg.WorkDir + "/" + cfg.StateFile
	err = os.WriteFile(outFile, jsonStr, 0644)
	if err != nil {
		return fmt.Errorf("cant write to out file %s, err: %w", outFile, err)
	}
	return nil
}

func extByName(name string) string {
	names := strings.Split(name, ".")
	if len(name) < 2 {
		return ""
	}
	return strings.ToLower(names[len(names)-1])
}

func fileinfo2State(file os.FileInfo) State {
	return State{
		Name:     strings.ToLower(strings.Replace(file.Name(), ".md", "", 1)),
		FileName: file.Name(),
		ModTime:  file.ModTime(),
		Size:     file.Size(),
	}
}
