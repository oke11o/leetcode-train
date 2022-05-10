package _2xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/sliding-window-maximum/
239. Sliding Window Maximum
Hard
#sliding_window, #deque
*/
func maxSlidingWindow(nums []int, k int) []int {
	deq := []int{}

	var cleanDeque = func(i, k int) {
		// remove indexes of elements not from sliding window
		if len(deq) != 0 && deq[0] == i-k {
			deq = deq[1:] // removeFirst
		}
		// remove from deq indexes of all elements
		// which are smaller than current element nums[i]
		for len(deq) != 0 && nums[i] > nums[deq[len(deq)-1]] { // deq.getLast()
			deq = deq[:len(deq)-1] // deq.removeLast()
		}
	}

	n := len(nums)
	if n*k == 0 {
		return []int{}
	}
	if k == 1 {
		return nums
	}

	maxIdx := 0
	for i := 0; i < k; i++ {
		cleanDeque(i, k)
		deq = append(deq, i)
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}
	output := make([]int, n-k+1)
	output[0] = nums[maxIdx]

	for i := k; i < n; i++ {
		cleanDeque(i, k)
		deq = append(deq, i)
		output[i-k+1] = nums[deq[0]] // deq.getFirst()
	}

	return output
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_maxSlidingWindow(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 3, 1, 2, 0, 5},
				k:    3,
			},
			want: []int{3, 3, 2, 5},
		},
		{
			name: "",
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "",
			args: args{
				nums: []int{8515, -3727, 5132, -6574, 5694, -9902, 9728, 9273, 4241, 9562, 5332, 7427, -9676, 6168, -9868, 9065, 9899, 7746, 3598, -3049, -7158, -1114, 2314, -9437, 54, -2872, -7914, -7147, 7334, -2529, -6288, -5125, 2767, -1376, 701, -1838, -4440, -8708, -8318, 789, -1319, -1992, 7939, 6125, 2635, 5413, 1705, -3327, -6058, -9787, -888, 8216, -4577, 7170, 2072, 5893, 8007, -265, 4503, 2036, 2422, 2789, 7767, 5928, 5410, 3843, -8397, -964, 8907, 9194, 5365, 9902, 2835, -2259, 8042, -3859, 9086, -9014, 380, 6786, -9503, -5941, 8298, -2160, 2297, 966, -2074, 1858, -7323, -2376, 2635, 4222, 1729, -6852, 9136, -8557, 662, -323, 4868, 9102, 8549, 3097, -5547, 7148, 5334, 1337, 5429, 8740, -3248, 7904, -7244, -6292, -5867, -2113, -5864, -6653, -5088, 6307, -9618, 9396, -2873, 6846, -6963, -9472, 8660, 5363, 8284, -5971, -8142, -2910, 8143, 8227, 2224, -6589, 6277, -3653, 1125, 3347, 7481, 5383, -8800, 984, 8276, -9796, 6676, 3835, -9870, 41, 1805, -2002, -6747, 6284, -3489, 5037, 7577, 36, -7440, 4773, -6140, -3767, -676, 8383, -6539, 5718, -6391, 5785, 6592, -6045, 5488, 835, -7716, 6612, 2039, 8807, -2809, 1398, 5097, -2201, 5345, 6372, 828, 9928, 6049, -2003, 3906, 2562, 429, 8770, 6257, 7429, -2122, -1644, -2975, -2873, -8823, -363, -4462, -7744, 8890, 3688, 1253, 5080, 3022, 9922, -7660, -1956, -8855, -6030, 3864, -8518, -8555, -3286, 1914, -1387, 8709, 1381, -956, -297, 1667, -2041, -9415, -3319, -4391, -9494, 9070, -1149, 4810, 6641, -8348, -6491, 8965, -3421, 9704, -4945, -326, 6437, 9657, -5745, -9369, -4178, 3601, 3415, 3510, 996, 5076, -6340, 9303, 1776, 6171, 2736, 266, -5351, 7685, 1322, 7766, -8830, -370, 7035, -2149, 470, -6246, 879, 3462, -1806, -7043, -7214, 6280, 5054, -3839, -6287, -3666, 5425, 5766, -201, 9793, 2713, 9375, -4813, 5456, 7138, 1109, -3933, 1466, -280, 903, 8422, -5896, -8448, -1052, 3876, 7263, 3565, 8457, 342, -656, -9711, 7213, 4318, -136, 5747, 2181, 5653, -1907, -4349, 9522, -5317, 6971, -6543, -9474, -8809, -9767, 7127, 4981, -4670, -142, 233, 1223, 4431, 4110, 6924, -9306, -2973, 1764, -5512, 5064, 3901, 390, -5218, -110, 4433, 899, 9449, 1631, -7321, -1366, -9099, 355, -174, 2632, -1649, -5562, -3244, -1566, 1575, -383, 8437, 4077, -4768, 3270, 6257, -5905, -1990, -2987, -7308, 918, 2900, -7966, 9239, -7214, -5333, 900, 802, -6043, 7905, 4289, -6867, 1013, -8818, 1538, -4135, 4411, -7907, 9169},
				k:    253,
			},
			want: []int{9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928, 9928},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}