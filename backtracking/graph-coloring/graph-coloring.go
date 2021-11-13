package graph_coloring

type Graph [][]int

var result []string

// https://www.youtube.com/watch?v=052VkKhIaQ4
//
func graphColoring(in Graph) []string {

	for _, c := range "RGB" {
		backtrack(in, 0, c, make([]rune, len(in)), "")
	}

	return nil
}

func backtrack(in Graph, idx int, curCol rune, curState []rune, current string) {
	current += string(curCol)
	if len(current) == len(in) {
		result = append(result, current)
		return
	}

	for i := idx; i < len(in); i++ {
		if !boundFunc(in, i, curCol, curState) {
			continue
		}
		//curState[i] = cur
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
