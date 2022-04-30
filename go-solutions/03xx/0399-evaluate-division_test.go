package _3xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/evaluate-division/
399. Evaluate Division
Medium

Мои замечания.
Тут интересный кейс, когда ищем деление на само себя. То есть пара `a`/`a`.
Можно было проверять данный кейс отдельно. Но тут так же работает корректно.
Сперва, находится пара a/b = 2. А потом это умножается b/a = 0.5.
*/
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i, v := range equations {
		score := values[i]
		if _, ok := graph[v[0]]; !ok {
			graph[v[0]] = make(map[string]float64)
		}
		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = make(map[string]float64)
		}
		graph[v[0]][v[1]] = score
		graph[v[1]][v[0]] = 1 / score
	}

	var backtrack func(graph map[string]map[string]float64, currNode, targetNode string, accProduct float64, visited map[string]struct{}) float64
	backtrack = func(graph map[string]map[string]float64, currNode, targetNode string, accProduct float64, visited map[string]struct{}) float64 {
		// mark the visit
		visited[currNode] = struct{}{}
		ret := -1.0

		neighbors := graph[currNode]
		if score, ok := neighbors[targetNode]; ok {
			ret = accProduct * score
		} else {
			for nextNode, score := range neighbors {
				if _, isVisited := visited[nextNode]; isVisited {
					continue
				}
				ret = backtrack(graph, nextNode, targetNode, accProduct*score, visited)
				if ret != -1.0 {
					break
				}
			}
		}

		// unmark the visit, for the next backtracking
		delete(visited, currNode)

		return ret
	}

	result := make([]float64, len(queries))
	for i, query := range queries {
		result[i] = -1.0
		if _, ok := graph[query[0]]; !ok {
			continue
		}
		if _, ok := graph[query[1]]; !ok {
			continue
		}
		visited := make(map[string]struct{})
		result[i] = backtrack(graph, query[0], query[1], 1, visited)
	}

	return result
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}},
				values:    []float64{2.0, 3.0},
				queries:   [][]string{{"a", "a"}, {"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			want: []float64{1.0, 6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			name: "",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
				values:    []float64{1.5, 2.5, 5.0},
				queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			},
			want: []float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			name: "",
			args: args{
				equations: [][]string{{"a", "b"}},
				values:    []float64{0.5},
				queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			},
			want: []float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}
