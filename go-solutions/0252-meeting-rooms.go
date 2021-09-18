package go_solutions

import "sort"

type Interval struct {
	Start, End int
}

func canAttendMeetings(intervals []*Interval) bool {
	if len(intervals) < 2 {
		return true
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1].End > intervals[i].Start {
			return false
		}
	}
	return true
}
