package main

import (
	"fmt"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6}
	s := array[1:2:4]
	fmt.Println(s, len(s), cap(s))
	a := s[1:]
	fmt.Println(a, len(a), cap(a))
}
