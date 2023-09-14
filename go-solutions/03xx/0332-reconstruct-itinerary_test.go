package _3xx

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

/*
*
https://leetcode.com/problems/reconstruct-itinerary/
332. Reconstruct Itinerary
Hard
Topics
Companies
You are given a list of airline tickets where tickets[i] = [fromi, toi] represent the departure and the arrival airports of one flight. Reconstruct the itinerary in order and return it.

All of the tickets belong to a man who departs from "JFK", thus, the itinerary must begin with "JFK". If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.

For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
You may assume all tickets form at least one valid itinerary. You must use all the tickets once and only once.
*/
func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	for _, t := range tickets {
		from := t[0]
		to := t[1]
		_, ok := graph[from]
		ch := []string{to}
		if !ok {
			graph[from] = ch
		} else {
			ch = append(ch, graph[from]...)
			graph[from] = ch
		}
	}

	st := []string{"JFK"}
	result := make([]string, 0)

	for len(st) > 0 {
		lastST := st[len(st)-1]
		if len(graph[lastST]) > 0 {
			st = append(st, graph[lastST][len(graph[lastST])-1])
			graph[lastST] = graph[lastST][:len(graph[lastST])-1]
		} else {
			st = st[:len(st)-1]
			result = append(result, lastST)
		}
	}

	l := len(result) - 1
	for i := 0; i < l/2; i++ {
		result[i], result[l-i] = result[l-i], result[i]
	}

	return result
}

func findItinerary___(tickets [][]string) []string {
	type node struct {
		key      string
		children []*node
		parent   *node
		weight   int
	}
	//build graph
	graph := map[string]*node{}
	for _, ticket := range tickets {
		from, ok := graph[ticket[0]]
		if !ok {
			from = &node{key: ticket[0]}
			graph[ticket[0]] = from
		}
		to, ok := graph[ticket[1]]
		if !ok {
			to = &node{key: ticket[1]}
			graph[ticket[1]] = to
		}
		from.children = append(from.children, to)
		to.parent = from
	}

	//weight
	length := len(graph)
	for _, n := range graph {
		v := n
		for i := 0; i < length; i++ {
			v = n.parent
			if v != nil {
				v.weight++
			} else {
				break
			}
		}
	}

	// sort
	sortedGraph := make([]*node, 0, len(graph))
	for _, v := range graph {
		sortedGraph = append(sortedGraph, v)
	}
	sort.Slice(sortedGraph, func(i, j int) bool { return (sortedGraph[i].weight) < (sortedGraph[j].weight) })

	// make result
	result := make([]string, len(sortedGraph))
	for i := 0; i < len(result); i++ {
		result[i] = sortedGraph[i].key
	}

	return result
}

func Test_findItinerary(t *testing.T) {
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{
			name:    "",
			tickets: [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			want:    []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name:    "",
			tickets: [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}},
			want:    []string{"JFK", "SFO", "ATL", "JFK", "ATL", "SFO"},
			//want:    []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findItinerary(tt.tickets)
			require.Equal(t, tt.want, got)
		})
	}
}
