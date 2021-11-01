package easy

func moveZerosToLeft(in []int) []int {
	rightPos := len(in) - 1
	for i := rightPos; i >= 0; i-- {
		if in[i] != 0 {
			in[rightPos], in[i] = in[i], in[rightPos]
			rightPos--
		}
	}
	return in
}
