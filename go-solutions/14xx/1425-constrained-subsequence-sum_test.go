package _4xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// I'd like to create Priority Queue for item [2]int.
// Where first element is priority and second element is value
type item interface {
	Val() int
	Less(item) bool
}
type MyPriorityQueue[T item] []T

func (h *MyPriorityQueue[T]) Parent(i int) int {
	return (i - 1) / 2
}
func (h *MyPriorityQueue[T]) LeftChild(i int) int {
	return 2*i + 1
}
func (h *MyPriorityQueue[T]) RightChild(i int) int {
	return 2*i + 2
}
func (h *MyPriorityQueue[T]) up(i int) {
	parent := h.Parent(i)
	for i > 0 && (*h)[parent].Less((*h)[i]) {
		(*h)[parent], (*h)[i] = (*h)[i], (*h)[parent]
		i = parent
		parent = h.Parent(i)
	}
}
func (h *MyPriorityQueue[T]) down(i int) {
	size := len((*h))
	maxIndex := i
	for {
		l := h.LeftChild(i)
		if l < size && (*h)[maxIndex].Less((*h)[l]) {
			maxIndex = l
		}
		r := h.RightChild(i)
		if r < size && (*h)[maxIndex].Less((*h)[r]) {
			maxIndex = r
		}
		if i != maxIndex {
			(*h)[maxIndex], (*h)[i] = (*h)[i], (*h)[maxIndex]
			i = maxIndex
		} else {
			break
		}
	}
}

func (h *MyPriorityQueue[T]) Insert(v T) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *MyPriorityQueue[T]) ExtractMax() T {
	result := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.down(0)

	return result
}

// 1425. Constrained Subsequence Sum
// Hard
// https://leetcode.com/problems/constrained-subsequence-sum

/*
Кратко: к максимуму из последних  k элементов прибавляем текущее

	   10   2   -10   5   20
	0  10  12     2  17   37

	    -1   -2    -3
	0   -1   -2    -3
*/
func constrainedSubsetSum_dp(nums []int, k int) int {
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	minInt := -1 << (intSize - 1)     // MinInt32 or MinInt64 depending on intSize.

	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums)+1)
	findKMax := func(from int) int { // к максимуму из последних  k элементов прибавляем текущее
		result := minInt
		for j := 0; j < k; j++ {
			i := from - j
			if i < 0 {
				return result
			}
			if result < dp[i] {
				result = dp[i]
			}
		}
		return result
	}
	for i := 0; i < len(nums); i++ {
		m := findKMax(i)
		dp[i+1] = m + nums[i]
	}
	result := minInt
	for i := 1; i < len(dp); i++ {
		if dp[i] > result {
			result = dp[i]
		}
	}

	return result
}

func Test_constrainedSubsetSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Explanation: The subsequence is [10, 2, 5, 20].",
			nums: []int{10, 2, -10, 5, 20},
			k:    2,
			want: 37,
		},
		{
			name: "Explanation: The subsequence must be non-empty, so we choose the largest number.",
			nums: []int{-1, -2, -3},
			k:    1,
			want: -1,
		},
		{
			name: "Explanation: The subsequence is [10, -2, -5, 20].",
			nums: []int{10, -2, -10, -5, 20},
			k:    2,
			want: 23,
		},
		{
			name: "",
			nums: []int{-5266, 4019, 7336, -3681, -5767},
			//  nums: []int{-5266,-1247, 6089, -3681, -5767},
			k:    2,
			want: 11355,
		},
		{
			name: "",
			nums: []int{0, 3217, -4023, -4138, -683, 6455, -3621, 9242, 4015, -3790},
			//  nums: []int{-8269,3217,-4023,-4138,-683,6455,-3621,9242,4015,-3790},   3217+ -4023+ -4138+ -683+ 6455+ -3621+ 9242+ 4015 = 10464  // -5627
			k:    1,
			want: 16091,
		},
		{
			name: "",
			nums: []int{-8269, 3217, -4023, -4138, -683, 6455, -3621, 9242, 4015, -3790},
			//  nums: []int{-8269,3217,-4023,-4138,-683,6455,-3621,9242,4015,-3790},
			k:    1,
			want: 16091,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constrainedSubsetSum(tt.nums, tt.k)
			require.Equal(t, tt.want, got)
		})
	}
}

type el [2]int

func (e el) Val() int {
	return e[0]
}
func (e el) Less(v item) bool {
	return e.Val() < v.Val()
}

func constrainedSubsetSum(nums []int, k int) int {
	maxF := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	heap := MyPriorityQueue[el]{}
	heap.Insert([2]int{nums[0], 0})
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		for i-heap[0][1] > k {
			heap.ExtractMax()
		}
		curr := maxF(0, heap[0][0]) + nums[i]
		ans = maxF(ans, curr)
		heap.Insert([2]int{curr, i})
	}

	return ans
}

// ===========

/*
GPT solution


Объясни, как работает PriorityQueue для следующего примера на Java?
class Solution {
    public int constrainedSubsetSum(int[] nums, int k) {
        PriorityQueue<int[]> heap = new PriorityQueue<>((a, b) -> {
            return b[0] - a[0];
        });

        heap.add(new int[] {nums[0], 0});
        int ans = nums[0];

        for (int i = 1; i < nums.length; i++) {
            while (i - heap.peek()[1] > k) {
                heap.remove();
            }

            int curr = Math.max(0, heap.peek()[0]) + nums[i];
            ans = Math.max(ans, curr);
            heap.add(new int[] {curr, i});
        }

        return ans;
    }
}


Данный пример решает задачу на нахождение максимальной суммы подряд идущих элементов в массиве, с учетом ограничения на размер подпоследовательности.

PriorityQueue в данном случае используется как макс-куча, которая хранит пары вида <сумма подмассива, индекс последнего элемента подмассива>. Компаратор установлен так, что элементы с наибольшей суммой имеют наивысший приоритет и находятся на вершине кучи.

Далее в цикле мы пробегаем каждый элемент (nums[i]) массива, начиная от второго:

Во время каждого шага цикла мы убираем элементы из кучи, пока не найдет наибольший элемент, который удовлетворяет условию (i - index) <= k, где i - текущий индекс, index - индекс подмассива, хранящегося в куче. Это гарантирует, что длина подмассива не будет превышать k.
Затем мы вычисляем текущую сумму подмассива, сравнивая текущий элемент массива и сумму элементов на вершине кучи. Если сумма верхушки кучи меньше нуля, мы избавляемся от нее, так как это уменьшает общую сумму.
Найденная сумма сравнивается с прежним максимумом и обновляет его при необходимости.
Наконец, мы добавляем текущую сумму и индекс в кучу.
Мы продолжаем этот процесс, пока не проходим все элементы массива. По окончанию цикла ans будет содержать искомое значение максимальной суммы подмассива.
Таким образом, данный код применяет технику скользящего окна и макс-кучи вместе, чтобы эффективно решить задачу на нахождение максимальной суммы подмассива с ограничением на его длину.
*/
