package _8xx

func peakIndexInMountainArray(arr []int) int {
	left := -1
	right := len(arr)
	for (right - left) > 1 {
		idx := (right + left) / 2
		if arr[idx] < arr[idx+1] {
			left = idx
		} else {
			right = idx
		}
	}
	return right
}
