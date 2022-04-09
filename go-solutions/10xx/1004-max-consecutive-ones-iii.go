package _0xx

/**
https://leetcode.com/problems/max-consecutive-ones-iii/
1004. Max Consecutive Ones III
Medium
#sliding_window
*/
func longestOnes(nums []int, k int) int {
	var ans int

	zeroes := 0
	units := 0

	first := 0
	second := 0
	//L_
	//[1,0,1,1,0]
	//F  ^
	//z=1
	//u=1
	//a=1
	for first < len(nums) {
		val1 := nums[first] // 0
		first++             // 2
		units++             // 2
		if val1 == 0 {
			zeroes++ //1
		}
		for zeroes > k { //false
			val2 := nums[second]
			second++
			units--
			if val2 == 0 {
				zeroes--
			}
		}
		if ans < units {
			ans = units
		}
	}

	return ans
}
