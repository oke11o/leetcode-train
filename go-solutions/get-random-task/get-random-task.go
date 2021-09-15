package main

import (
	"math/rand"
	"time"
)

// Get random task
// Выбираем задачу. Учитываются только задачи с минимальных score
// Когда выполняю задачу, то поднимаю score++
func main() {
	rand.Seed(time.Now().UnixMicro())
	m := map[string]int{
		"0046": 0,
		"0047": 0,
		"0054": 0,
		"0056": 0,
		"0073": 0,
		"0078": 0,
		"0079": 0,
		"0090": 1,
		"0098": 0,
		"0100": 0,
		"0141": 0,
		"0153": 0,
		"0198": 0,
		"0202": 0,
		"0206": 0,
		"0234": 0,
		"238":  1,
		"287":  0,
		"0340": 0,
		"442":  0,
		"0448": 0,
		"0784": 0,
		"0876": 0,
		"0904": 0,
		"0917": 0,
		"0978": 0,
		"1189": 0,
	}
	min := 9999
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	l := make([]string, 0, len(m))
	for k, v := range m {
		if v == min {
			l = append(l, k)
		}
	}
	i := rand.Intn(len(l))
	println(l[i])
}
