package amazon

func deviations(s string) int32 {
	var max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	var min = func(i, j int) int {
		if i < j {
			return i
		}
		return j
	}

	if len(s) == 0 {
		return 0
	}
	index_dict := make([][]int, 26)
	for i, v := range s {
		index_dict[v-'a'] = append(index_dict[v-'a'], i)
	}

	var ans int
	for i := 0; i < 25; i++ {
		for j := i + 1; j < 26; j++ {
			list1, list2 := index_dict[i], index_dict[j]
			a, b := 0, 0
			count := 0 // the number of first letter -the number of second letter in [0:min(list1[a],list2[b])+1]
			//  stores the max and min of count along the way

			max_diff, min_diff := 0, 1<<31-1
			for a < len(list1) || b < len(list2) {
				if a == len(list1) {
					b += 1
					count -= 1
				} else if b == len(list2) {
					a += 1
					count += 1
				} else if list1[a] < list2[b] { //: # means we encountered first letter first
					count += 1
					a += 1
				} else {
					count -= 1
					b += 1
				}
				max_diff = max(max_diff, count)
				min_diff = min(min_diff, count)
			}
			ans = max(ans, max_diff-min_diff)
		}
	}

	return int32(ans)
}
func deviations_(s string) int32 {
	var ans int32
	first := 0
	last := 0
	// _
	// bbacccabab
	//   ^
	buf := make(map[int32]int, 26)
	for first < len(s) {
		l1 := int32(s[first]) // b
		buf[l1]++             // {b:2,a:1}
		first++               // 2
		for len(buf) > 2 {    //false
			l2 := int32(s[last])
			buf[l2]--
			last++
			if buf[l2] == 0 {
				delete(buf, l2)
			}
		}
		// find freq
		max := int(-1)
		min := int(27)
		for _, v := range buf {
			if max < v {
				max = v
			}
			if min > v {
				min = v
			}
		}
		tmp := int32(max - min)
		if ans < tmp {
			ans = tmp
		}
	}

	return ans
}
