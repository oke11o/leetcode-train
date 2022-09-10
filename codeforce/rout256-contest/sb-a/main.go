package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var err error
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	_, err = fmt.Fscan(in, &cnt)
	panicErr(err)

	for i := 0; i < cnt; i++ {
		var n, m int
		_, err = fmt.Fscan(in, &n, &m)
		panicErr(err)
		fmt.Println(out, n+m)
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
