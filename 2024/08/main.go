package main

import (
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/08/input.txt", solvePart1, solvePart2)
}

type node struct {
	idx int
	row int
	col int
	fq  byte
}

func solvePart1(data string) (int, error) {
	bytes := []byte(data)
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	result := 0
	nodes := []node{}
	antinodes := make(map[int]bool, len(bytes))
	for i, byte := range bytes {
		if byte != '.' && byte != '\n' {
			row, col := indexToCoord(max_cols, i)
			nodes = append(nodes, node{i, row, col, byte})
		}
	}
	for _, current := range nodes {
		for _, other := range nodes {
			if other.fq != current.fq || other.idx == current.idx {
				continue
			}
			anti_row, anti_col := getAntiNodeCoord(current.row, current.col, other.row, other.col)
			if outOfBounds(max_rows, max_cols, anti_row, anti_col) {
				continue
			}
			anti_idx := coordToIndex(max_cols, anti_row, anti_col)
			if !antinodes[anti_idx] {
				result += 1
				antinodes[anti_idx] = true
			}
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	bytes := []byte(data)
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	result := 0
	nodes := []node{}
	antinodes := make(map[int]bool, len(bytes))
	for i, byte := range bytes {
		if byte != '.' && byte != '\n' {
			row, col := indexToCoord(max_cols, i)
			nodes = append(nodes, node{i, row, col, byte})
		}
	}
	for _, current := range nodes {
		for _, other := range nodes {
			if other.fq != current.fq || other.idx == current.idx {
				continue
			}
			if !antinodes[other.idx] {
				result += 1
				antinodes[other.idx] = true
			}
			cur_row := current.row
			cur_col := current.col
			oth_row := other.row
			oth_col := other.col
			for true {
				anti_row, anti_col := getAntiNodeCoord(cur_row, cur_col, oth_row, oth_col)
				if outOfBounds(max_rows, max_cols, anti_row, anti_col) {
					break
				}
				anti_idx := coordToIndex(max_cols, anti_row, anti_col)
				if !antinodes[anti_idx] {
					result += 1
					antinodes[anti_idx] = true
				}
				cur_row = oth_row
				cur_col = oth_col
				oth_row = anti_row
				oth_col = anti_col
			}
		}
	}
	return result, nil
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

func getAntiNodeCoord(cur_row int, cur_col int, oth_row int, oth_col int) (int, int) {
	return oth_row + (oth_row - cur_row), oth_col + (oth_col - cur_col)
}
