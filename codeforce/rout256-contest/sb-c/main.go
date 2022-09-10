package main

import (
	"bufio"
	"fmt"
	"math"
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

		for _, r := range problemC(arr) {
			fmt.Print(r[0], r[1], "\r\n")
		}
		if i != cnt-1 {
			fmt.Print("\r\n")
		}
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func problemC(arr []int) [][2]int {
	cache := make(map[int]struct{}, len(arr)/2)
	result := [][2]int{}
	from := 0
	for n := len(arr) / 2; n > 0; n-- {
		min := 99999999
		second := from + 1
		for i := second; i < len(arr); i++ {
			if _, ok := cache[i]; ok {
				continue
			}
			a := int(math.Abs(float64(arr[from] - arr[i])))
			if a < min {
				min = a
				second = i
			}
			if a == 0 {
				break
			}
		}
		result = append(result, [2]int{from + 1, second + 1})
		cache[second] = struct{}{}
		from++
		for {
			if _, ok := cache[from]; ok {
				from++
			} else {
				break
			}
		}
	}

	return result
}
