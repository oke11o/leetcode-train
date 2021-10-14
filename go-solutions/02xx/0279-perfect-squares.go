package _2xx

const MaxInt32 = 1<<31 - 1

var cache = make(map[int]int)

// 0279. Perfect Squares
func numSquares(n int) int {
	if n < 4 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	mins := []int{}
	i := 1
	for {
		ii := i * i
		if ii > n {
			break
		} else if ii == n {
			return 1
		}
		mins = append(mins, 1+numSquares(n-ii))
		i++
	}

	res := min(mins)
	cache[n] = res
	return res
}

func min(nums []int) int {
	res := int(MaxInt32)
	for _, v := range nums {
		if res > v {
			res = v
		}
	}
	return res
}
