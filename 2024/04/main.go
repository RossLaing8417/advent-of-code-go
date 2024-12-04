package main

import (
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/04/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	lines := strings.Split(data, "\n")
	result := 0
	for row, line := range lines {
		for col := range line {
			result += countXMAS(lines, row, col)
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	lines := strings.Split(data, "\n")
	result := 0
	for row := range lines[2:] {
		for col := range lines[row][2:] {
			result += countMAS(lines, row, col)
		}
	}
	return result, nil
}

func countXMAS(lines []string, row int, col int) int {
	result := 0
	deltas := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, delta := range deltas {
		if matchXMAS(lines, row, delta[0], col, delta[1]) {
			result += 1
		}
	}
	return result
}

func matchXMAS(lines []string, row int, drow int, col int, dcol int) bool {
	if row+drow*3 < 0 || row+drow*3 >= len(lines) || col+dcol*3 < 0 || col+dcol*3 >= len(lines[row]) {
		return false
	}
	for _, char := range "XMAS" {
		if lines[row][col] != byte(char) {
			return false
		}
		row += drow
		col += dcol
	}
	return true
}

func countMAS(lines []string, row int, col int) int {
	result := 0
	deltas := [][2][4]int{
		{{2, 2, -1, -1}, {2, 0, -1, 1}},
		{{2, 2, -1, -1}, {0, 2, 1, -1}},
		{{0, 0, 1, 1}, {2, 0, -1, 1}},
		{{0, 0, 1, 1}, {0, 2, 1, -1}},
	}
	for _, delta := range deltas {
		first := matchMAS(lines, row+delta[0][0], delta[0][2], col+delta[0][1], delta[0][3])
		second := matchMAS(lines, row+delta[1][0], delta[1][2], col+delta[1][1], delta[1][3])
		if first && second {
			result += 1
		}
	}
	return result
}

func matchMAS(lines []string, row int, drow int, col int, dcol int) bool {
	for _, char := range "MAS" {
		if lines[row][col] != byte(char) {
			return false
		}
		row += drow
		col += dcol
	}
	return true
}
