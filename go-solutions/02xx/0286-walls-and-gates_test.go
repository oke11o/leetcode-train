package _2xx

import (
	"reflect"
	"testing"
)

/**
https://leetcode.com/problems/walls-and-gates/
286. Walls and Gates
Medium
#BFS, #matrix, #array

Мое пояснение:
Обратить внимание, что мы добавили все GATE. А потом оттуда пошли обычным BFS с обычным условием.
Хотя тут как раз условие важно. Но об этом после.
Так как это BFS, то мы просто на каждой итерации добавляем в нашу очередь следующий уровень.
А так как мы идем от ворот, то по сути нам это и надо.
То есть мы сперва обрабатываем уровень 0.
Потом от него добавляем уровень 1. (В условии это rooms[r][c] != EMPTY)
И так далее.
И граничное условие - мы не идем глубже в листы, если уже установили верное значение. И еще раз повторю, нам это и надо было.
*/
func wallsAndGates(rooms [][]int) {
	const (
		EMPTY = 1<<31 - 1
		GATE  = 0
	)
	var directions = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	m := len(rooms)
	if m == 0 {
		return
	}
	n := len(rooms[0])

	q := [][2]int{}
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if rooms[row][col] == GATE {
				q = append(q, [2]int{row, col})
			}
		}
	}

	for len(q) != 0 {
		point := q[0]
		q = q[1:]
		row := point[0]
		col := point[1]
		for _, dir := range directions {
			r := row + dir[0]
			c := col + dir[1]
			if r < 0 || c < 0 || r >= m || c >= n {
				continue
			}
			if rooms[r][c] != EMPTY {
				continue
			}
			rooms[r][c] = rooms[row][col] + 1
			q = append(q, [2]int{r, c})
		}
	}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_wallsAndGates(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				rooms: [][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}},
			},
			want: [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}},
		},
		{
			name: "",
			args: args{
				rooms: [][]int{{-1}},
			},
			want: [][]int{{-1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallsAndGates(tt.args.rooms)
			if !reflect.DeepEqual(tt.args.rooms, tt.want) {
				t.Errorf("DeepEqual() = %v, want %v", tt.args.rooms, tt.want)
			}
		})
	}
}
