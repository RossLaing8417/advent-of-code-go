package main

import (
	"slices"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/06/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	bytes := []byte(data)
	index := slices.Index(bytes, '^')
	dir := bytes[index]
	result := 0
	for true {
		row, col := indexToCoord(max_cols, index)
		row, col = move(row, col, dir)
		if outOfBounds(max_rows, max_cols, row, col) {
			bytes[index] = 'X'
			result += 1
			break
		}
		new_index := coordToIndex(max_cols, row, col)
		if bytes[new_index] == '#' {
			dir = rotate(dir)
		} else {
			if bytes[index] != 'X' {
				bytes[index] = 'X'
				result += 1
			}
			index = new_index
		}
	}
	return result, nil
}

type node struct {
	row int
	col int
	dir byte
}

var DEBUG bool = false

func solvePart2(data string) (int, error) {
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	bytes := []byte(data)
	index := slices.Index(bytes, '^')
	dir := bytes[index]
	visited := make([]node, 0)
	result := 0
	for true {
		row, col := indexToCoord(max_cols, index)
		chk_row, chk_col := move(row, col, dir)
		if outOfBounds(max_rows, max_cols, chk_row, chk_col) {
			bytes[index] = 'X'
			break
		}
		chk_index := coordToIndex(max_cols, chk_row, chk_col)
		if bytes[chk_index] == '#' {
			dir = rotate(dir)
		} else {
			bytes[index] = 'X'
			visited = append(visited, node{row, col, dir})
			if bytes[chk_index] != 'X' {
				tmp := bytes[chk_index]
				bytes[chk_index] = '#'
				if loopExists(bytes, visited, max_rows, max_cols, row, col, rotate(dir)) {
					result += 1
				}
				bytes[chk_index] = tmp
			}
			index = chk_index
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

func move(row int, col int, dir byte) (int, int) {
	switch dir {
	case '^':
		row -= 1
	case '>':
		col += 1
	case 'v':
		row += 1
	case '<':
		col -= 1
	}
	return row, col
}

func rotate(dir byte) byte {
	switch dir {
	case '^':
		dir = '>'
	case '>':
		dir = 'v'
	case 'v':
		dir = '<'
	case '<':
		dir = '^'
	}
	return dir
}

func outOfBounds(max_rows int, max_cols int, row int, col int) bool {
	return row < 0 || row >= max_rows || col < 0 || col >= max_cols
}

func hasObsticlesOnPath(bytes []byte, max_rows int, max_cols int, row int, col int, dir byte) bool {
	for !outOfBounds(max_rows, max_cols, row, col) {
		if bytes[coordToIndex(max_cols, row, col)] == '#' {
			return true
		}
		row, col = move(row, col, dir)
	}
	return false
}

func loopExists(bytes []byte, visited []node, max_rows int, max_cols int, row int, col int, dir byte) bool {
	nodes := append([]node{}, visited...)
	for slices.Index(nodes, node{row, col, dir}) == -1 {
		nodes = append(nodes, node{row, col, dir})
		chk_row, chk_col := move(row, col, dir)
		chk_index := coordToIndex(max_cols, chk_row, chk_col)
		if outOfBounds(max_rows, max_cols, chk_row, chk_col) {
			return false
		}
		if bytes[chk_index] == '#' {
			dir = rotate(dir)
		} else {
			row = chk_row
			col = chk_col
		}
	}
	return true
}
