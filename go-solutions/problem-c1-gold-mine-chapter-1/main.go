package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultFileName = "sample_input.txt"

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		_, _ = fmt.Fprintf(os.Stdout, "unexpected exit: %s", err)
		return
	}
	_, _ = fmt.Fprint(os.Stdout, "\n\nDONE\n")
}

func run(args []string, stdout *os.File, stderr *os.File) error {

	in, err := readInput(getInputFilename(args))
	_ = in

	return err
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
				item.Pairs = append(item.Pairs, [2]int{ints[0], ints[1]})
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
