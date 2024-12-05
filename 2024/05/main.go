package main

import (
	"slices"
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/05/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	order_map := make(map[int][]int)
	order_rules := true
	result := 0
line_loop:
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			order_rules = false
			continue
		}
		if order_rules {
			pages := strings.Split(line, "|")
			l_value, err := strconv.Atoi(pages[0])
			if err != nil {
				return 0, err
			}
			r_value, err := strconv.Atoi(pages[1])
			if err != nil {
				return 0, err
			}
			order_map[l_value] = append(order_map[l_value], r_value)
		} else {
			pages := []int{}
			for _, page := range strings.Split(line, ",") {
				num, err := strconv.Atoi(page)
				if err != nil {
					return 0, err
				}
				list, valid := order_map[num]
				if valid {
					for _, item := range list {
						if slices.IndexFunc(pages, func(p int) bool { return p == item }) != -1 {
							continue line_loop
						}
					}
				}
				pages = append(pages, num)
			}
			result += pages[len(pages)/2]
		}
	}

	return result, nil
}

func solvePart2(data string) (int, error) {
	order_map := make(map[int][]int)
	order_rules := true
	result := 0
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			order_rules = false
			continue
		}
		if order_rules {
			pages := strings.Split(line, "|")
			l_value, err := strconv.Atoi(pages[0])
			if err != nil {
				return 0, err
			}
			r_value, err := strconv.Atoi(pages[1])
			if err != nil {
				return 0, err
			}
			order_map[l_value] = append(order_map[l_value], r_value)
		} else {
			pages := []int{}
			valid_line := true
			for _, page := range strings.Split(line, ",") {
				num, err := strconv.Atoi(page)
				if err != nil {
					return 0, err
				}
				if valid_line {
					list, valid := order_map[num]
					if valid {
						for _, item := range list {
							if slices.IndexFunc(pages, func(p int) bool { return p == item }) != -1 {
								valid_line = false
								break
							}
						}
					}
				}
				pages = append(pages, num)
			}
			if !valid_line {
				slices.SortFunc(pages, func(a int, b int) int {
					return slices.IndexFunc(order_map[a], func(item int) bool { return item == b })
				})
				result += pages[len(pages)/2]
			}
		}
	}

	return result, nil
}
