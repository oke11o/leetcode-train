package _7xx

// 0744. Find Smallest Letter Greater Than Target
func nextGreatestLetter(letters []byte, target byte) byte {
	left := -1
	right := len(letters)
	for (right - left) > 1 {
		idx := (right + left) / 2
		if letters[idx] > target {
			right = idx
		} else {
			left = idx
		}
	}

	return letters[(left+1)%len(letters)]
}
