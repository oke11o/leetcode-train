package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const defaultFileName = "gold_mine_chapter_1_input.txt"

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "unexpected exit: %s", err)
		return
	}
	_, _ = fmt.Fprint(os.Stdout, "\n\nDONE\n")
}

func run(args []string, stdout *os.File, stderr *os.File) error {
	in, err := readInput(getInputFilename(args))
	if err != nil {
		return err
	}
	var out strings.Builder
	for i, inn := range in.Inputs {
		o := solve(inn)
		out.WriteString(fmt.Sprintf("Case #%d: %d\n", i+1, o))
	}
	return ioutil.WriteFile("gold_mine_chapter_1_result.txt", []byte(out.String()), 0644)
}

func solve(input inputItem) int {
	n := input.T
	adj := make(map[int][]int, n)
	dp := make([]int, n)
	seen := make([]bool, n)
	for i, v := range input.Pairs {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
		dp[i] = 0
		seen[i] = false
	}

	stack := make([]int, 1, 1)

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		if !seen[node] {
			seen[node] = true
			for _, ad := range adj[node] {
				if !seen[ad] {
					stack = append(stack, ad)
				}
			}
		} else {
			stack = stack[: len(stack)-1 : len(stack)-1]
			seen[node] = false
			child_max := 0
			for _, ad := range adj[node] {
				if !seen[ad] {
					if child_max < dp[ad] {
						child_max = dp[ad]
					}
				}
			}
			dp[node] = child_max + input.Values[node]
		}
	}

	if len(adj[0]) <= 1 {
		return dp[0]
	}
	var max1, max2 int

	for _, ad := range adj[0] {
		if dp[ad] > max2 {
			if dp[ad] > max1 {
				max2 = max1
				max1 = dp[ad]
			} else {
				max2 = dp[ad]
			}
		}
	}
	return max1 + max2 + input.Values[0]
}

func getInputFilename(args []string) string {
	if len(args) < 2 {
		return defaultFileName
	}
	return args[1]
}

type inputItem struct {
	T      int
	Values []int
	Pairs  [][2]int
}

type commonInput struct {
	T      int
	Inputs []inputItem
}

func readInput(filename string) (commonInput, error) {
	result := commonInput{}
	file, err := os.Open(filename)
	if err != nil {
		return result, fmt.Errorf("Cant open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	j := 0
	for scanner.Scan() {
		i++
		v := scanner.Text()
		t, err := strconv.Atoi(v)
		if err != nil {
			return result, nil
		}
		if i == 1 {
			result.T = t
			continue
		}

		item := inputItem{T: t}
		for j = t; j > 0; j-- {
			if !scanner.Scan() {
				return result, errors.New("invalid input unknown")
			}
			v := scanner.Text()
			vv := strings.Split(v, " ")
			ints := make([]int, 0, len(vv))
			for _, vvv := range vv {
				if value, err := strconv.Atoi(vvv); err == nil {
					ints = append(ints, value)
				} else {
					return result, nil
				}
			}
			if j == t {
				item.Values = ints
			} else if len(ints) == 2 {
				item.Pairs = append(item.Pairs, [2]int{ints[0] - 1, ints[1] - 1})
			} else {
				return result, fmt.Errorf("invalid input: invalid int pair: %s", v)
			}
		}
		result.Inputs = append(result.Inputs, item)
	}

	if err := scanner.Err(); err != nil {
		return result, fmt.Errorf("Scanner err: %w", err)
	}

	return result, nil
}
