package _5xx

import (
	"container/heap"
	"math"
	"sort"
	"testing"
)

/*
*
https://leetcode.com/problems/min-cost-to-connect-all-points/
1584. Min Cost to Connect All Points
Medium
*/
func minCostConnectPoints(points [][]int) int {
	type edge struct {
		weight   float64
		curNode  int
		nextNode int
	}
	n := len(points)
	allEdges := []edge{}

	// Storing all edges of our complete graph.
	for currNode := 0; currNode < n; currNode++ {
		for nextNode := currNode + 1; nextNode < n; nextNode++ {
			weight := math.Abs(float64(points[currNode][0]-points[nextNode][0])) +
				math.Abs(float64(points[currNode][1]-points[nextNode][1]))

			allEdges = append(allEdges, edge{weight, currNode, nextNode})
		}
	}

	// Sort all edges in increasing order.
	sort.Slice(allEdges, func(i, j int) bool {
		return allEdges[i].weight < allEdges[j].weight
	})

	uf := NewUnionFind(n)
	mstCost := float64(0)
	edgesUsed := 0

	for i := 0; i < len(allEdges) && edgesUsed < n-1; i++ {
		edge := allEdges[i]

		if uf.connect(edge.curNode, edge.nextNode) {
			mstCost += edge.weight
			edgesUsed++
		}
	}

	return int(mstCost)
}

type UnionFind struct {
	rating  []int
	parents []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		rating:  make([]int, n),
		parents: make([]int, n),
	}
	for i := range uf.parents {
		uf.parents[i] = i
	}
	return uf
}

// connect return flag. Is already connected
func (uf *UnionFind) connect(a, b int) bool {
	parentA := uf.find(a)
	parentB := uf.find(b)
	if parentA == parentB {
		return false
	}
	if uf.rating[parentA] == uf.rating[parentB] {
		uf.parents[parentA] = parentB
		uf.rating[parentB]++
	} else if uf.rating[parentA] < uf.rating[parentB] {
		uf.parents[parentA] = parentB
	} else {
		uf.parents[parentB] = parentA
	}

	return true
}

func (uf *UnionFind) find(a int) int {
	if uf.parents[a] == a {
		return a
	}
	return uf.find(uf.parents[a])
}

type item struct {
	node   int
	weight int
	index  int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].weight > pq[j].weight
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	i := x.(*item)
	i.index = n
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) update(item *item, node int, weight int) {
	item.node = node
	item.weight = weight
	heap.Fix(pq, item.index)
}

func minCostConnectPoints2(points [][]int) int {

	//heap.Fix()
	return 0
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_minCostConnectPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			},
			want: 20,
		},
		{
			name: "",
			args: args{
				points: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints2(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
