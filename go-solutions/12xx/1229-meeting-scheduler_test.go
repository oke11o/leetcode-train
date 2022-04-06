package _2xx

import (
	"reflect"
	"sort"
	"testing"
)

/**
https://leetcode.com/problems/meeting-scheduler/
1229. Meeting Scheduler
Medium
*/
func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	i := 0
	j := 0
	sort.Slice(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.Slice(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})
	//slots1 = [[0,2]],
	//		     ^
	//slots2 = [[1,3]]
	//			 ^
	for i < len(slots1) && j < len(slots2) {
		start1 := slots1[i][0]
		end1 := slots1[i][1]
		start2 := slots2[j][0]
		end2 := slots2[j][1]
		if end1 < start2 { //  3 < 2 no
			i++
		} else if end2 < start1 { // 3 < 0 no
			j++
		} else {
			end := end1     // 2
			if end > end2 { //2>3
				end = end2 // 2
			}
			start := start1     //0
			if start < start2 { //0<1
				start = start2 // 1
			}
			intesect := end - start   //2-1
			if intesect >= duration { //no
				return []int{start, start + duration}
			}
			if end1 < end2 { //50 < 15
				i++
			} else {
				j++
			}
		}
	}
	return []int{}
}

/*********************************/
/************* TESTS *************/
/*********************************/
func Test_minAvailableDuration(t *testing.T) {
	type args struct {
		slots1   [][]int
		slots2   [][]int
		duration int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				slots1:   [][]int{{10, 50}, {60, 120}, {140, 210}},
				slots2:   [][]int{{0, 15}, {60, 70}},
				duration: 8,
			},
			want: []int{60, 68},
		},
		{
			name: "",
			args: args{
				slots1:   [][]int{{10, 50}, {60, 120}, {140, 210}},
				slots2:   [][]int{{0, 15}, {60, 70}},
				duration: 12,
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				slots1:   [][]int{{0, 2}},
				slots2:   [][]int{{1, 3}},
				duration: 1,
			},
			want: []int{1, 2},
		},
		{
			name: "",
			args: args{
				slots1: [][]int{
					{216397070, 363167701},
					{98730764, 158208909},
					{441003187, 466254040},
					{558239978, 678368334},
					{683942980, 717766451},
				},
				slots2:   [][]int{{50490609, 222653186}, {512711631, 670791418}, {730229023, 802410205}, {812553104, 891266775}, {230032010, 399152578}},
				duration: 456085,
			},
			want: []int{98730764, 99186849},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAvailableDuration(tt.args.slots1, tt.args.slots2, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAvailableDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
