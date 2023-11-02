package main

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1     // MaxInt32 or MaxInt64 depending on intSize.
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

//func (h *MyPriorityQueue[T]) Remove(i int) {
//	(*h)[i][0] = MaxInt
//	h.up(i)
//	h.ExtractMax()
//}
