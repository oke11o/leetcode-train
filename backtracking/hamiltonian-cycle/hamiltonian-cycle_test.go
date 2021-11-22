package hamiltonian_cycle

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_hamiltonianCycle(t *testing.T) {
	tests := []struct {
		name  string
		graph Graph
		want  [][]int
	}{
		{
			name:  "",
			graph: Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			want:  [][]int{{0, 1, 2, 3, 4, 0}, {0, 1, 4, 3, 2, 0}, {0, 2, 1, 3, 4, 0}, {0, 2, 3, 1, 4, 0}, {0, 2, 3, 4, 1, 0}, {0, 4, 1, 3, 2, 0}, {0, 4, 3, 1, 2, 0}, {0, 4, 3, 2, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hamiltonianCycle(tt.graph)
			require.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_boundFunc(t *testing.T) {
	tests := []struct {
		name         string
		graph        Graph
		position     int
		vertexVal    int
		currentState []int
		want         bool
	}{
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     1,
			vertexVal:    1,
			currentState: []int{0, 0, 0, 0, 0},
			want:         true,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     2,
			vertexVal:    1,
			currentState: []int{0, 1, 0, 0, 0},
			want:         false,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     2,
			vertexVal:    2,
			currentState: []int{0, 1, 0, 0, 0},
			want:         true,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     3,
			vertexVal:    0,
			currentState: []int{0, 1, 2, 0, 0},
			want:         false,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     3,
			vertexVal:    1,
			currentState: []int{0, 1, 2, 0, 0},
			want:         false,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     3,
			vertexVal:    2,
			currentState: []int{0, 1, 2, 0, 0},
			want:         false,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     3,
			vertexVal:    3,
			currentState: []int{0, 1, 2, 0, 0},
			want:         true,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     3,
			vertexVal:    4,
			currentState: []int{0, 1, 2, 0, 0},
			want:         false,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     4,
			vertexVal:    4,
			currentState: []int{0, 1, 2, 3, 0},
			want:         true,
		},
		{
			name:         "",
			graph:        Graph{0: {1, 2, 4}, 1: {0, 2, 3, 4}, 2: {0, 1, 3}, 3: {1, 2, 4}, 4: {0, 1, 3}},
			position:     2,
			vertexVal:    4,
			currentState: []int{0, 2, 0, 0, 0},
			want:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := boundFunc(tt.graph, tt.position, tt.vertexVal, tt.currentState)
			require.Equal(t, tt.want, got)
		})
	}
}
