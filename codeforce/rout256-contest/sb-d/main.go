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
		var inCnt int
		_, err = fmt.Fscan(in, &inCnt)
		panicErr(err)

		arr := make([]int, 0, inCnt)
		for j := 0; j < inCnt; j++ {
			var a int
			_, err = fmt.Fscan(in, &a)
			panicErr(err)
			arr = append(arr, a)
		}

		d := problemD(arr)
		fmt.Print(d, "\r\n")
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func problemD(arr []int) string {
	cache := make(map[int]struct{})
	for i, a := range arr {
		if i == 0 {
			cache[a] = struct{}{}
			continue
		}
		if arr[i-1] == a {
			continue
		}
		if _, ok := cache[a]; ok {
			return "NO"
		}
		cache[a] = struct{}{}
	}
	return "YES"
}
