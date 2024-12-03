package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/02/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	var result int = 0
line_loop:
	for _, line := range strings.Split(data, "\n") {
		strs := strings.Split(line, " ")
		nums := make([]int, len(strs))
		for i, num := range strs {
			tmp, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			nums[i] = tmp
		}
		neg := nums[0+1]-nums[0] < 0
		for i := 1; i < len(nums); i += 1 {
			tmp := nums[i] - nums[i-1]
			if (absInt(tmp) > 3 || absInt(tmp) == 0) || (tmp < 0 && !neg) || (tmp > 0 && neg) {
				continue line_loop
			}
		}
		result += 1
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	var result int = 0
	for _, line := range strings.Split(data, "\n") {
		strs := strings.Split(line, " ")
		nums := make([]int, len(strs))
		for i, num := range strs {
			tmp, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			nums[i] = tmp
		}
		safe := isSafe(nums)
		if !safe {
			tmp_nums := make([]int, len(nums)-1)
			for x := 0; x < len(nums); x += 1 {
				if x > 0 {
					copy(tmp_nums[0:], nums[0:x])
				}
				if x < len(nums)-1 {
					copy(tmp_nums[x:], nums[x+1:])
				}
				safe = isSafe(tmp_nums)
				if safe {
					break
				}
			}
		}
		if safe {
			result += 1
		}
	}
	return result, nil
}

func isSafe(nums []int) bool {
	neg := nums[0+1]-nums[0] < 0
	for i := 1; i < len(nums); i += 1 {
		tmp := nums[i] - nums[i-1]
		if (absInt(tmp) > 3 || absInt(tmp) == 0) || (tmp < 0 && !neg) || (tmp > 0 && neg) {
			return false
		}
	}
	return true
}

func absInt(num int) int {
	return int(math.Abs(float64(num)))
}
