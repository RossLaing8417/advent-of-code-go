package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/01/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	left_nums := make([]int, strings.Count(data, "\n")+1)
	right_nums := make([]int, len(left_nums))
	for i, line := range strings.Split(data, "\n") {
		var (
			left  int
			right int
		)
		if _, err := fmt.Sscanf(line, "%d   %d", &left, &right); err != nil {
			return 0, err
		}
		println(left, right)
		left_nums[i] = left
		right_nums[i] = right
	}
	slices.Sort(left_nums)
	slices.Sort(right_nums)
	result := 0
	for i := range left_nums {
		result += int(math.Abs(float64(left_nums[i]) - float64(right_nums[i])))
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	left_nums := make([]int, strings.Count(data, "\n")+1)
	right_nums := make([]int, len(left_nums))
	for i, line := range strings.Split(data, "\n") {
		var (
			left  int
			right int
		)
		if _, err := fmt.Sscanf(line, "%d   %d", &left, &right); err != nil {
			return 0, err
		}
		println(left, right)
		left_nums[i] = left
		right_nums[i] = right
	}
	slices.Sort(left_nums)
	slices.Sort(right_nums)
	result := 0
	for _, left := range left_nums {
		for _, right := range right_nums {
			if left == right {
				result += left
			}
		}
	}
	return result, nil
}
