package go_solutions

func maxTurbulenceSize(arr []int) int {
	if len(arr) == 1 {
		return 1
	}

	const (
		up   = "up"
		down = "down"
	)

	var isOpposite = func(prev, cur string) bool {
		//if prev == "" && cur != "" {
		//	return true
		//}
		if prev == up && cur == down {
			return true
		}
		if prev == down && cur == up {
			return true
		}
		return false
	}
	var getDir = func(i, j int) string {
		if i < j {
			return up
		}
		if i > j {
			return down
		}
		return ""
	}

	result := 1
	counter := 1
	prevDir := ""
	for i := 1; i < len(arr); i++ {
		curDir := getDir(arr[i-1], arr[i])
		if isOpposite(prevDir, curDir) {
			counter++
		} else {
			counter++
			if result < counter {
				result = counter
			}
			counter = 1
		}
		prevDir = curDir
	}
	if result < counter {
		result = counter
	}
	return result
}
