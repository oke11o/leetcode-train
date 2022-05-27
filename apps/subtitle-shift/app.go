package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const dateSeparator = " --> "

type Srt struct {
	Num   string // "\ufeff1"
	Date  string // 00:02:58,225 --> 00:03:00,561
	From  int64
	To    int64
	Lines []string
	// We shouldn't be here
	// when it gets dark.
}

func run(args []string, stdout *os.File, stderr *os.File) error {
	cfg, err := readConfig()
	if err != nil {
		return fmt.Errorf("cant read cfg: %w", err)
	}
	outFile, err := os.OpenFile(cfg.ToFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cant create outFile: %w", err)
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)
	_ = writer

	inFile, err := os.Open(cfg.FromFile)
	if err != nil {
		return fmt.Errorf("cant open inFile: %w", err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	srt := Srt{}
	state := 0
	i := -1
	for scanner.Scan() {
		i++
		txt := scanner.Text()
		switch state {
		case 0:
			srt.Num = txt
			state = 1
		case 1:
			srt.Date = txt
			srt.From, srt.To, err = parseDate(txt)
			if err != nil {
				return fmt.Errorf("line: %d, error: %w", i, err)
			}
			state = 2
		case 2:
			if txt == "" {
				state = 0
				srt.From += cfg.Increment
				srt.To += cfg.Increment
				_, err := writer.WriteString(fmt.Sprintf("%s\n%s\n", srt.Num, dates2date(srt.From, srt.To)))
				if err != nil {
					return fmt.Errorf("invalid OUT on line `%d`, %w", i, err)
				}
				for _, t := range srt.Lines {
					_, err := writer.WriteString(t + "\n")
					if err != nil {
						return fmt.Errorf("invalid OUT on line `%d`, %w", i, err)
					}
				}
				_, err = writer.WriteString("\n")
				if err != nil {
					return fmt.Errorf("invalid OUT on line `%d`, %w", i, err)
				}
				srt.Lines = []string{}
				err = writer.Flush()
				if err != nil {
					return fmt.Errorf("invalid Flush on line `%d`, %w", i, err)
				}
			} else {
				srt.Lines = append(srt.Lines, txt)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func parseDate(date string) (int64, int64, error) {
	parse := func(dt string) (int64, error) {
		dts := strings.Split(dt, ",")
		if len(dts) != 2 {
			return 0, fmt.Errorf("`%s` not found `,`", dt)
		}
		result, err := strconv.ParseInt(dts[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("`%s` cant parse miliseconds, %w", dt, err)
		}
		seconds := strings.Split(dts[0], ":")
		if len(seconds) != 3 {
			return 0, fmt.Errorf("`%s` cant parse seconds", dts[0])
		}
		s, err := strconv.ParseInt(seconds[2], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("`%s` cant parse seconds, %w", dt, err)
		}
		result += s * 1000
		s, err = strconv.ParseInt(seconds[1], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("`%s` cant parse minutes, %w", dt, err)
		}
		result += s * 1000 * 60
		s, err = strconv.ParseInt(seconds[0], 10, 64)
		if err != nil {
			return 0, fmt.Errorf("`%s` cant parse hours, %w", dt, err)
		}
		result += s * 1000 * 60 * 60

		return result, nil
	}
	strs := strings.Split(date, dateSeparator)
	if len(strs) != 2 {
		return 0, 0, fmt.Errorf("invalid parseDate format: %s", date)
	}

	from, err := parse(strs[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parseDate format: %s, %w", date, err)
	}
	to, err := parse(strs[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid parseDate format: %s, %w", date, err)
	}
	_, _ = from, to

	return from, to, nil
}

func dates2date(from, to int64) string {
	convert := func(val int64) string {
		var h, m, s, mili int64
		mili = val % 1000
		val /= 1000
		s = val % 60
		val /= 60
		m = val % 60
		h = val / 60

		return fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, mili)
	}
	f := convert(from)
	t := convert(to)
	return f + dateSeparator + t
}
