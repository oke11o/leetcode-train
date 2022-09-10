package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

		arr := make([]string, 0, inCnt)
		for j := 0; j < inCnt; j++ {
			var a string
			_, err = fmt.Fscan(in, &a)
			panicErr(err)
			arr = append(arr, a)
		}

		d := problemE(arr)
		fmt.Print(d, "\r\n")
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func problemE(arr []string) string {
	le, err := parseAndValidateList(arr)
	if err != nil {
		return "NO"
	}
	if intersectionsExists(le) {
		return "NO"
	}
	return "YES"
}

func intersectionsExists(in [][2]int) bool {
	if len(in) < 2 {
		return false
	}
	sort.Slice(in, func(i, j int) bool {
		return in[i][0] < in[j][0]
	})
	for i := 1; i < len(in); i++ {
		if in[i][0] <= in[i-1][1] {
			return true
		}
	}
	return false
}

func parseAndValidateList(arr []string) ([][2]int, error) {
	result := make([][2]int, 0, len(arr))
	for _, a := range arr {
		aa := strings.Split(a, "-")
		if len(aa) != 2 {
			return nil, fmt.Errorf("Wrong input `%s`", a)
		}
		b0, err := parseAndValidate(aa[0])
		if err != nil {
			return nil, err
		}
		b1, err := parseAndValidate(aa[1])
		if err != nil {
			return nil, err
		}
		if b0 > b1 {
			return nil, fmt.Errorf("Wrong input `%s`", a)
		}
		result = append(result, [2]int{b0, b1})
	}
	return result, nil
}

func parseAndValidate(s string) (int, error) {
	ss := strings.Split(s, ":")
	if len(ss) != 3 {
		return 0, fmt.Errorf("Wrong input `%s`", s)
	}
	sec, err := strconv.Atoi(ss[2])
	if err != nil {
		return 0, err
	}
	if sec > 59 {
		return 0, fmt.Errorf("Wrong input `%s`", s)
	}
	min, err := strconv.Atoi(ss[1])
	if min > 59 {
		return 0, fmt.Errorf("Wrong input `%s`", s)
	}
	hour, err := strconv.Atoi(ss[0])
	if hour > 23 {
		return 0, fmt.Errorf("Wrong input `%s`", s)
	}
	return hour*60*60 + min*60 + sec, nil
}
