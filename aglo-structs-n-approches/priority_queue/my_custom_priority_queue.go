package main

const (
	intSize = 32 << (^uint(0) >> 63) // 32 or 64
	MaxInt  = 1<<(intSize-1) - 1     // MaxInt32 or MaxInt64 depending on intSize.
)

// I'd like to create Priority Queue for item [2]int.
// Where first element is priority and second element is value

type MyPriorityQueue [][2]int

func (h *MyPriorityQueue) Parent(i int) int {
	return (i - 1) / 2
}
func (h *MyPriorityQueue) LeftChild(i int) int {
	return 2*i + 1
}
func (h *MyPriorityQueue) RightChild(i int) int {
	return 2*i + 2
}
func (h MyPriorityQueue) up(i int) {
	parent := h.Parent(i)
	for i > 0 && h[parent][0] < h[i][0] {
		h[parent], h[i] = h[i], h[parent]
		i = parent
		parent = h.Parent(i)
	}
}
func (h MyPriorityQueue) down(i int) {
	size := len(h)
	maxIndex := i
	for {
		l := h.LeftChild(i)
		if l < size && h[l][0] > h[maxIndex][0] {
			maxIndex = l
		}
		r := h.RightChild(i)
		if r < size && h[r][0] > h[maxIndex][0] {
			maxIndex = r
		}
		if i != maxIndex {
			h[maxIndex], h[i] = h[i], h[maxIndex]
			i = maxIndex
		} else {
			break
		}
	}
}

func (h *MyPriorityQueue) Insert(v [2]int) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *MyPriorityQueue) ExtractMax() [2]int {
	result := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.down(0)

	return result
}
func (h *MyPriorityQueue) Remove(i int) {
	(*h)[i][0] = MaxInt
	h.up(i)
	h.ExtractMax()
}
