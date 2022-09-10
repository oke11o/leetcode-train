package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

		fmt.Println(problemB(arr))
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func problemB(arr []int) int {
	sort.Ints(arr)
	var result int
	cnt := 0
	for i, x := range arr {
		if i == 0 {
			result += x
			cnt++
			continue
		}
		if arr[i-1] != x {
			cnt = 0
		}
		if cnt == 2 {
			cnt = 0
			continue
		}
		result += x
		cnt++
	}
	return result
}
