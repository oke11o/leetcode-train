package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	for i, v := range x { // длина слайса рассчитывается на первой итерации
		fmt.Printf("%d) %s [l=%d,c=%d]", i, v, len(x), cap(x))
		// fmt.Println(i, v, len(x))
		// fmt.Println(x[i]) // panic: runtime error: index out of range [3] with length 3
		// получается что, по индексу мы обратиться не можем, но range проитерируется до конца, т.е. выйдет за len(x)
		if v == "a" {
			x = append(x[:i], x[i+1:]...)
		}
		fmt.Print("\n")
	}
}
