package graph_coloring

type Graph [][]int

var result []string

// https://www.youtube.com/watch?v=052VkKhIaQ4
//
func graphColoring(in Graph) []string {

	backtrack(in, 0, make([]rune, len(in)), "")

	return result
}

// R = 82
// G = 71
func backtrack(in Graph, idx int, curState []rune, current string) {
	if len(current) == len(in) {
		result = append(result, current)
		return
	}

	for _, c := range "RGB" {
		if !boundFunc(in, idx, c, curState) {
			continue
		}
		newState := make([]rune, len(curState))
		copy(newState, curState)
		newState[idx] = c
		backtrack(in, idx+1, newState, current+string(c))
	}
	a := 1
	_ = a
}

func boundFunc(in Graph, idx int, curColor rune, curState []rune) bool {
	if idx > len(in)-1 {
		panic("invalid input Graph")
	}
	for _, subling := range in[idx] {
		if subling > len(curState)-1 {
			panic("invalid current state")
		}
		if curState[subling] == curColor {
			return false
		}
	}
	return true
}
