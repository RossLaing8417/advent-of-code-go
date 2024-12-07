package main

import (
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

type Operator int

const (
	ADD Operator = iota
	MUL
	CON
)

func main() {
	aoc.Execute("2024/07/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	for _, line := range strings.Split(data, "\n") {
		split := strings.Split(line, ": ")
		expected, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, nil
		}
		split = strings.Split(split[1], " ")
		nums := make([]int, len(split))
		for i, n := range split {
			nums[i], err = strconv.Atoi(n)
			if err != nil {
				return 0, nil
			}
		}
		if eval(expected, nums, []Operator{ADD, MUL}) {
			result += expected
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	for _, line := range strings.Split(data, "\n") {
		split := strings.Split(line, ": ")
		expected, err := strconv.Atoi(split[0])
		if err != nil {
			return 0, nil
		}
		split = strings.Split(split[1], " ")
		nums := make([]int, len(split))
		for i, n := range split {
			nums[i], err = strconv.Atoi(n)
			if err != nil {
				return 0, nil
			}
		}
		if eval(expected, nums, []Operator{ADD, MUL, CON}) {
			result += expected
		}
	}
	return result, nil
}

func eval(current int, nums []int, operators []Operator) bool {
	if len(nums) == 1 {
		return current == nums[0]
	}
	index := len(nums) - 1
	for _, operator := range operators {
		cur := current
		switch operator {
		case ADD:
			cur -= nums[index]
			if cur < 0 {
				continue
			}
		case MUL:
			if cur%nums[index] != 0 {
				continue
			}
			cur /= nums[index]
		case CON:
			tmp_cur := strconv.Itoa(cur)
			tmp_num := strconv.Itoa(nums[index])
			if len(tmp_num) >= len(tmp_cur) || !strings.HasSuffix(tmp_cur, tmp_num) {
				continue
			}
			c, err := strconv.Atoi(tmp_cur[:len(tmp_cur)-len(tmp_num)])
			if err != nil {
				panic(err)
			}
			cur = c
		}
		if eval(cur, nums[:index], operators) {
			return true
		}
	}
	return false
}
