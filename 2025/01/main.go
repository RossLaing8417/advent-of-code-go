package main

import (
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/01/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result, _, err := solve(data, 50)
	return result, err
}

func solvePart2(data string) (int, error) {
	_, result, err := solve(data, 50)
	return result, err
}

func solve(data string, start int) (int, int, error) {
	var value, count_zeros, count_passes, delta int = start, 0, 0, 1
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			continue
		}
		i, err := strconv.ParseInt(line[1:], 10, 0)
		if err != nil {
			return 0, 0, err
		}
		if line[0] == 'L' {
			delta = -1
		} else {
			delta = 1
		}
		for range int(i) {
			value += delta
			if value%100 == 0 {
				count_passes += 1
			}
		}
		if value%100 == 0 {
			count_zeros += 1
		}
	}
	return count_zeros, count_passes, nil
}
