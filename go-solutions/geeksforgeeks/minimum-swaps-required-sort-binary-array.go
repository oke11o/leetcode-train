package geeksforgeeks

/**
 * Time complexity: O(n^2)
 * Space complexity: O(n) - because we sorted twice!
 */
func findMinSwaps_bruteForce(input []int) int {
	result0 := 0
	input0 := make([]int, len(input))
	copy(input0, input)
	for i := 0; i < len(input0)-1; i++ {
		for j := 0; j < len(input0)-1; j++ {
			if input0[j] > input0[j+1] {
				input0[j], input0[j+1] = input0[j+1], input0[j]
				result0++
			}
		}
	}
	result1 := 0
	for i := 0; i < len(input)-1; i++ {
		for j := 0; j < len(input)-1; j++ {
			if input[j] < input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
				result1++
			}
		}
	}
	if result0 < result1 {
		return result0
	}
	return result1
}

/**
 * https://www.geeksforgeeks.org/minimum-swaps-required-sort-binary-array/
 * Minimum adjacent swaps required to Sort Binary array
 */
func findMinSwaps(input []int) int {
	unitsLeft := 0
	result := 0
	for _, val := range input {
		if val == 1 {
			unitsLeft++
		} else {
			result += unitsLeft
		}
	}

	unitsLeft = 0
	result1 := 0
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 1 {
			unitsLeft++
		} else {
			result1 += unitsLeft
		}
	}

	if result < result1 {
		return result
	}
	return result1
}
