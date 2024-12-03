package main

import (
	"fmt"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/03/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	l_value := 0
	r_value := 0
	result := 0
	for _, line := range strings.Split(data, "\n") {
		for i := 0; i < len(line)-4; i += 1 {
			substr := line[i:]
			if !strings.HasPrefix(substr, "mul(") {
				continue
			}
			r_paren := strings.Index(substr, ")")
			if r_paren == -1 {
				i += 4
				continue
			}
			_, err := fmt.Sscanf(substr, "mul(%d,%d)", &l_value, &r_value)
			if err != nil {
				continue
			}
			i = i + r_paren
			result += l_value * r_value
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	l_value := 0
	r_value := 0
	result := 0
	enabled := true
	for _, line := range strings.Split(data, "\n") {
		for i := 0; i < len(line)-4; i += 1 {
			substr := line[i:]
			if enabled && strings.HasPrefix(substr, "don't()") {
				enabled = false
				i += 6
				continue
			}
			if !enabled && strings.HasPrefix(substr, "do()") {
				enabled = true
				i += 3
				continue
			}
			if !enabled || !strings.HasPrefix(substr, "mul(") {
				continue
			}
			r_paren := strings.Index(substr, ")")
			if r_paren == -1 {
				i += 4
				continue
			}
			_, err := fmt.Sscanf(substr, "mul(%d,%d)", &l_value, &r_value)
			if err != nil {
				continue
			}
			i = i + r_paren
			result += l_value * r_value
		}
	}
	return result, nil
}
