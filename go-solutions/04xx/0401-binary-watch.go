package _4xx

import "fmt"

var readBinaryWatchResult []string

/**
 * 0401. Binary Watch
 * Easy
 * Backtracking, Bit Manipulation
 * https://leetcode.com/problems/binary-watch/
 */
func readBinaryWatch(turnedOn int) []string {
	readBinaryWatchResult = []string{}

	readBinaryWatchBacktrack(turnedOn, 10, 0)

	return readBinaryWatchResult
}

func readBinaryWatchBacktrack(turnedOn int, level int, current int16) {
	if turnedOn > level {
		return
	}
	if level == 0 {
		hours := current / 64
		minutes := current % 64
		if hours > 11 {
			return
		}
		if minutes > 59 {
			return
		}
		if turnedOn < 0 {
			return
		}
		if turnedOn == 0 {
			readBinaryWatchResult = append(readBinaryWatchResult, watch2string(current))
		}
		return
	}
	level--
	current = current << 1
	for _, v := range []int{1, 0} {
		if v == 1 {
			readBinaryWatchBacktrack(turnedOn-1, level, current+1)
		} else {
			readBinaryWatchBacktrack(turnedOn, level, current)
		}
	}
}

func watch2string(current int16) string {
	return fmt.Sprintf("%d:%02d", current/64, current%64)
}
