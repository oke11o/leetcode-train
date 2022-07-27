package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "unexpected exit: %s \n", err)
	}
	fmt.Println("DONE! Thank you for using! See you")
}

func run(args []string, stdout *os.File, stderr *os.File) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("cant load .env file: %w", err)
	}
	cfg := initConfig()

	var arg string
	if len(args) > 1 {
		arg = args[1]
	}
	switch arg {
	case "build-state":
		err = buildState(cfg)
	case "scan-file":
		err = scanFile(cfg, false)
	default:
		err = scanFile(cfg, true)
	}
	return err
}
