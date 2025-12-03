package main

import (
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/03/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	return solve(data, 2)
}

func solvePart2(data string) (int, error) {
	return solve(data, 12)
}

func solve(data string, size int) (int, error) {
	result := 0
	for _, l := range strings.Split(data, "\n") {
		line := []byte(l)
		on := ""
		prev := -1
		for itr := range size {
			cur := prev + 1
			end := len(line) - (size - (itr + 1))
			for i := cur + 1; i < end; i += 1 {
				if line[i] > line[cur] {
					cur = i
				}
			}
			on += strconv.FormatInt(int64(line[cur]-'0'), 10)
			prev = cur
		}
		jolts, err := strconv.ParseUint(on, 10, 64)
		if err != nil {
			return 0, err
		}
		result += int(jolts)
	}
	return result, nil
}
