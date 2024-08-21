package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5} //l=5 c=5

	b := a[2:] //345 l:3 c=3
	b = append(b, 8)

	b[0] = 9

	fmt.Println(a)
}
