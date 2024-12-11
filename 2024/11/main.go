package main

import (
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/11/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	nums := []int{}
	for _, line := range strings.Split(data, "\n") {
		for _, num := range strings.Split(line, " ") {
			value, err := strconv.Atoi(num)
			if err != nil {
				return 0, nil
			}
			nums = append(nums, value)
			value, err = solve(value, 25)
			if err != nil {
				return 0, err
			}
			result += value
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	for _, line := range strings.Split(data, "\n") {
		for _, num := range strings.Split(line, " ") {
			value, err := strconv.Atoi(num)
			if err != nil {
				return 0, err
			}
			value, err = solve(value, 75)
			if err != nil {
				return 0, err
			}
			result += value
		}
	}
	return result, nil
}

type fnCache struct {
	num    int
	count  int
	result int
}

var cache []fnCache = []fnCache{}

func solve(num int, count int) (int, error) {
	if count == 0 {
		return 1, nil
	}
	for _, c := range cache {
		if c.num == num && c.count == count {
			return c.result, nil
		}
	}
	result := 0
	if num == 0 {
		tmp, err := solve(1, count-1)
		if err != nil {
			return 0, err
		}
		result += tmp
	} else {
		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			tmp, err := strconv.Atoi(str[0 : len(str)/2])
			if err != nil {
				return 0, err
			}
			tmp, err = solve(tmp, count-1)
			if err != nil {
				return 0, err
			}
			result += tmp
			tmp, err = strconv.Atoi(str[len(str)/2:])
			if err != nil {
				return 0, err
			}
			tmp, err = solve(tmp, count-1)
			result += tmp
		} else {
			tmp, err := solve(num*2024, count-1)
			if err != nil {
				return 0, nil
			}
			result += tmp
		}
	}
	cache = append(cache, fnCache{num, count, result})
	return result, nil
}
