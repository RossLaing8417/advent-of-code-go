package main

import (
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/06/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	lines := strings.Split(data, "\n")
	operands, operators := [][]int{}, []string{}
	for i, line := range lines {
		entries := strings.Fields(line)
		if i < len(lines)-1 {
			nums := []int{}
			for _, entry := range entries {
				if num, err := strconv.ParseInt(entry, 10, 0); err != nil {
					return 0, err
				} else {
					nums = append(nums, int(num))
				}
			}
			operands = append(operands, nums)
		} else {
			operators = entries
		}
	}
	for i, operator := range operators {
		res := operands[0][i]
		for j := range operands[1:] {
			switch operator {
			case "+":
				res += operands[j+1][i]
			case "*":
				res *= operands[j+1][i]
			}
		}
		result += res
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	lines := strings.Split(data, "\n")
	line_len := 0
	for _, line := range lines {
		line_len = max(line_len, len(line))
	}
	for i, line := range lines {
		for len(line) < line_len {
			line += " "
		}
		lines[i] = line
	}
	nums := []int{}
	for i := line_len - 1; i >= 0; i -= 1 {
		tmp := ""
		for j := 0; j < len(lines)-1; j += 1 {
			if lines[j][i] != ' ' {
				tmp += string(lines[j][i])
			}
		}
		if tmp == "" {
			continue
		}
		if num, err := strconv.ParseInt(tmp, 10, 0); err != nil {
			return 0, err
		} else {
			nums = append(nums, int(num))
		}
		if lines[len(lines)-1][i] != ' ' {
			res := nums[0]
			for _, num := range nums[1:] {
				switch lines[len(lines)-1][i] {
				case '+':
					res += num
				case '*':
					res *= num
				}
			}
			nums = nums[:0]
			result += res
		}
	}
	return result, nil
}
