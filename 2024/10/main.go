package main

import (
	"slices"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/10/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	bytes := []byte(data)
	result := 0
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	for i, item := range bytes {
		if item == '0' {
			result += solve(bytes, max_rows, max_cols, i, false)
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	bytes := []byte(data)
	result := 0
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	for i, item := range bytes {
		if item == '0' {
			result += solve(bytes, max_rows, max_cols, i, true)
		}
	}
	return result, nil
}

func solve(bytes []byte, max_rows int, max_cols int, start int, part2 bool) int {
	queue := []int{start}
	index := 0
	deltas := [][]int{
		{-1, 0},
		{0, -1},
		{1, 0},
		{0, 1},
	}
	ends := []int{}
	for len(queue) > 0 {
		index = queue[0]
		queue = queue[1:]
		row, col := indexToCoord(max_cols, index)
		for _, delta := range deltas {
			new_row := row + delta[0]
			new_col := col + delta[1]
			if outOfBounds(max_rows, max_cols, new_row, new_col) {
				continue
			}
			new_index := coordToIndex(max_cols, new_row, new_col)
			if bytes[new_index] == bytes[index]+1 {
				if bytes[new_index] == '9' {
					if part2 || slices.Index(ends, new_index) == -1 {
						ends = append(ends, new_index)
					}
				} else {
					if part2 || slices.Index(queue, new_index) == -1 {
						queue = append(queue, new_index)
					}
				}
			}
		}
	}
	return len(ends)
}

func indexToCoord(max_cols int, index int) (int, int) {
	row := index / (max_cols + 1)
	return row, index - row - (row * max_cols)
}

func coordToIndex(max_cols int, row int, col int) int {
	return (row * max_cols) + row + col
}

func outOfBounds(max_rows int, max_cols int, row int, col int) bool {
	return row < 0 || row >= max_rows || col < 0 || col >= max_cols
}
