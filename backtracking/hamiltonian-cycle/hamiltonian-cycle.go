package hamiltonian_cycle

type Graph [][]int

var result [][]int

/**
 * https://www.youtube.com/watch?v=dQr4wZCiJJ4&list=PLDN4rrl48XKpZkf03iYFl-O29szjTrs_O&index=67
 * 6.4 Hamiltonian Cycle - Backtracking
 */
func hamiltonianCycle(graph Graph) [][]int {
	result = make([][]int, 0)
	backtrack(graph, 1, make([]int, len(graph)))
	return result
}

func backtrack(graph Graph, position int, currentState []int) {
	if position == len(graph) {
		if hasPair(graph, currentState[len(currentState)-1], 0) {
			result = append(result, append(currentState, 0))
		}
		return
	}

	// Проходим по все нодам графа, проверям, можно ли в эту позицию использовать эту ноду
	for val := 1; val < len(graph); val++ {
		if boundFunc(graph, position, val, currentState) {
			tmp := make([]int, len(currentState))
			copy(tmp, currentState)
			tmp[position] = val
			// Если можно, идем к следующей позиции. И для нее так же проверим все ноды
			backtrack(graph, position+1, tmp)
		}
	}
}

func boundFunc(graph Graph, position, vertexVal int, currentState []int) bool {
	for i := position - 1; i >= 0; i-- {
		posVal := currentState[i]
		if vertexVal == posVal {
			return false
		}
	}
	return hasPair(graph, currentState[position-1], vertexVal)
}

func hasPair(graph Graph, vertex int, target int) bool {
	for _, v := range graph[vertex] {
		if v == target {
			return true
		}
	}
	return false
}
