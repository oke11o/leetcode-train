package _4xx

var readBinaryWatchResult []string

/**
 * 0401. Binary Watch
 * Easy
 * Backtracking, Bit Manipulation
 * https://leetcode.com/problems/binary-watch/
 */
func readBinaryWatch(turnedOn int) []string {
	readBinaryWatchResult = []string{}

	readBinaryWatchBacktrack(turnedOn)

	return readBinaryWatchResult
}

func readBinaryWatchBacktrack(turnedOn int) {

}
