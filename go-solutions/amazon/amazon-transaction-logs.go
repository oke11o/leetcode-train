package amazon

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'processLogs' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY logs
 *  2. INTEGER threshold
 */

// Amazon Transaction logs
func processLogs(logs []string, threshold int32) []string {
	m := make(map[string]int32)
	for _, log := range logs {
		l := strings.Split(log, " ")
		if l[0] == l[1] {
			m[l[0]]++
		} else {
			m[l[0]]++
			m[l[1]]++
		}
	}
	result := []string{}
	users := []int{}
	for u, c := range m {
		if c >= threshold {
			result = append(result, u)
			ui, _ := strconv.Atoi(u)
			users = append(users, ui)
		}
	}

	sort.Ints(users)
	result = make([]string, len(users))
	for i, u := range users {
		result[i] = strconv.Itoa(u)
	}

	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	logsCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var logs []string

	for i := 0; i < int(logsCount); i++ {
		logsItem := readLine(reader)
		logs = append(logs, logsItem)
	}

	thresholdTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	threshold := int32(thresholdTemp)

	result := processLogs(logs, threshold)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
