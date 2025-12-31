package main

import (
	"bytes"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/07/input.txt", solvePart1, solvePart2)
}

func solvePart1(sdata string) (int, error) {
	result := 0
	data := []byte(sdata)
	lines := bytes.Split(data, []byte("\n"))
	prev := lines[0]
	for _, line := range lines[1:] {
		for i := range prev {
			if prev[i] != '|' && prev[i] != 'S' {
				continue
			}
			if line[i] == '^' {
				line[i-1] = '|'
				line[i+1] = '|'
				result += 1
			} else {
				line[i] = '|'
			}
		}
		prev = line
	}
	return result, nil
}

// func solvePart2(sdata string) (int, error) {
// 	result := 0
// 	data := []byte(sdata)
// 	data = bytes.ReplaceAll(data, []byte("."), []byte{0})
// 	data = bytes.Replace(data, []byte("S"), []byte{1}, 1)
// 	lines := bytes.Split(data, []byte("\n"))
// 	prev := lines[0]
// 	for _, line := range lines[1:] {
// 		for i := range prev {
// 			if prev[i] == 0 || prev[i] == '^' {
// 				continue
// 			}
// 			if line[i] == '^' {
// 				line[i-1] += prev[i]
// 				line[i+1] += prev[i]
// 			} else {
// 				line[i] += prev[i]
// 			}
// 		}
// 		prev = line
// 	}
// 	for _, b := range lines[len(lines)-1] {
// 		result += int(b)
// 	}
// 	return result, nil
// }

// NOTE:
// Above works for example but not for input (gives 9411 which is way off)
// Tried programming it exactly as seen in a gif posted on reddit and it worked...
// https://www.reddit.com/r/adventofcode/comments/1pgnmou/2025_day_7_lets_visualize/

func solvePart2(sdata string) (int, error) {
	result := 0
	lines := strings.Split(sdata, "\n")
	nums := make([]int, len(lines))
	lidx, ridx := strings.Index(lines[0], "S"), strings.Index(lines[0], "S")
	nums[lidx] = 1
	for _, line := range lines {
		for i := lidx; i <= ridx; i += 1 {
			if line[i] == '^' {
				nums[i-1] += nums[i]
				nums[i+1] += nums[i]
				nums[i] = 0
				lidx = min(lidx, i-1)
				ridx = max(ridx, i+1)
			}
		}
	}
	for _, num := range nums {
		result += num
	}
	return result, nil
}
