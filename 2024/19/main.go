package main

import (
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/19/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	split := strings.Split(data, "\n\n")
	towels := strings.Split(split[0], ", ")
	result := 0
	for _, pattern := range strings.Split(split[1], "\n") {
		if canCreatePattern(pattern, towels) > 0 {
			result += 1
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	split := strings.Split(data, "\n\n")
	towels := strings.Split(split[0], ", ")
	result := 0
	for _, pattern := range strings.Split(split[1], "\n") {
		result += canCreatePattern(pattern, towels)
	}
	return result, nil
}

var func_cache map[string]int = make(map[string]int)

func canCreatePattern(pattern string, towels []string) int {
	if len(pattern) == 0 {
		// fmt.Printf("Success\n")
		return 1
	}
	if result, existing := func_cache[pattern]; existing {
		return result
	}
	// fmt.Printf("Pattern: %s\n", pattern)
	result := 0
	for _, t := range towels {
		if strings.HasPrefix(pattern, t) {
			result += canCreatePattern(pattern[len(t):], towels)
		}
	}
	func_cache[pattern] = result
	return result
}
