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
		if prevDir == "" && curDir == "" {
			counter = 0
		} else if isOpposite(prevDir, curDir) {
			counter++
		} else {
			if result < counter+1 {
				result = counter + 1
			}
			counter = 1
		}
		prevDir = curDir
	}
	if result < counter+1 {
		result = counter + 1
	}
	return result
}
