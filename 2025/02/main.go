package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/02/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	for _, ids := range strings.Split(data, ",") {
		var start, end int64 = 0, 0
		_, err := fmt.Sscanf(ids, "%d-%d", &start, &end)
		if err != nil {
			return 0, nil
		}
		for i := start; i <= end; i += 1 {
			str := strconv.FormatInt(i, 10)
			if len(str)%2 != 0 {
				continue
			}
			mid := len(str) / 2
			if str[:mid] == str[mid:] {
				result += int(i)
			}
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	for _, ids := range strings.Split(data, ",") {
		var start, end int64 = 0, 0
		_, err := fmt.Sscanf(ids, "%d-%d", &start, &end)
		if err != nil {
			return 0, nil
		}
		for i := start; i <= end; i += 1 {
			str := strconv.FormatInt(i, 10)
			slen := len(str)
			for x := 2; x <= slen; x += 1 {
				if slen%x != 0 {
					continue
				}
				idx := slen / x
				if strings.Count(str[idx:], str[:idx]) == x-1 {
					result += int(i)
					break
				}
			}
		}
	}
	return result, nil
}
