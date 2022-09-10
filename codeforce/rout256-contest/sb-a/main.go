package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var err error
	in := bufio.NewReader(os.Stdin)
	var cnt int
	_, err = fmt.Fscan(in, &cnt)
	panicErr(err)

	for i := 0; i < cnt; i++ {
		var a int
		_, err = fmt.Fscan(in, &a)
		panicErr(err)
		var b int
		_, err = fmt.Fscan(in, &b)

		fmt.Println(fn(a, b))
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fn(a, b int) int {
	return a + b
}
