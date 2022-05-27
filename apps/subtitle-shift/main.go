package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "unexpected exit: %s \n", err)
	}
	fmt.Println("DONE! Thank you for using! See you")
}
