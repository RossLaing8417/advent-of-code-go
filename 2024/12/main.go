package main

import (
	"slices"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/12/input.txt", solvePart1, solvePart2)
}

type node struct {
	row            int
	col            int
	adjacent_nodes []int
}

func solvePart1(data string) (int, error) {
	bytes := []byte(data)
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	result := 0
	nodes := map[int]node{}
	edges := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for index, char := range bytes {
		if char == '\n' {
			continue
		}
		row, col := indexToCoord(max_cols, index)
		current := node{row, col, []int{}}
		for _, delta := range edges {
			new_row, new_col := row+delta[0], col+delta[1]
			if outOfBounds(max_rows, max_cols, new_row, new_col) {
				continue
			}
			new_idx := coordToIndex(max_cols, new_row, new_col)
			if bytes[new_idx] == bytes[index] {
				current.adjacent_nodes = append(current.adjacent_nodes, new_idx)
			}
		}
		nodes[index] = current
	}
	visited := []int{}
	queue := []int{}
	for key := range nodes {
		if slices.Index(visited, key) != -1 {
			continue
		}
		queue = append(queue[:0], key)
		area := 0
		perimeter := 0
		for len(queue) > 0 {
			index := queue[0]
			visited = append(visited, index)
			current := nodes[index]
			queue = queue[1:]
			area += 1
			perimeter += 4 - len(current.adjacent_nodes)
			for _, j := range current.adjacent_nodes {
				if slices.Index(visited, j) != -1 || slices.Index(queue, j) != -1 {
					continue
				}
				queue = append(queue, j)
			}
		}
		result += area * perimeter
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	bytes := []byte(data)
	max_rows := strings.Count(data, "\n") + 1
	max_cols := strings.Index(data, "\n")
	result := 0
	nodes := map[int]node{}
	edges := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	corners := [][]int{{-1, 0, -1, -1, 0, -1}, {0, -1, 1, -1, 1, 0}, {1, 0, 1, 1, 0, 1}, {0, 1, -1, 1, -1, 0}}
	for index, char := range bytes {
		if char == '\n' {
			continue
		}
		row, col := indexToCoord(max_cols, index)
		current := node{row, col, []int{}}
		for _, delta := range edges {
			new_row, new_col := row+delta[0], col+delta[1]
			if outOfBounds(max_rows, max_cols, new_row, new_col) {
				continue
			}
			new_idx := coordToIndex(max_cols, new_row, new_col)
			if bytes[new_idx] == bytes[index] {
				current.adjacent_nodes = append(current.adjacent_nodes, new_idx)
			}
		}
		nodes[index] = current
	}
	visited := []int{}
	section := []int{}
	queue := []int{}
	for key := range nodes {
		if slices.Index(visited, key) != -1 {
			continue
		}
		queue = append(queue[:0], key)
		section = section[:0]
		for len(queue) != 0 {
			index := queue[0]
			queue = queue[1:]
			section = append(section, index)
			for _, adjacent := range nodes[index].adjacent_nodes {
				if slices.Index(queue, adjacent) == -1 && slices.Index(section, adjacent) == -1 {
					queue = append(queue, adjacent)
				}
			}
		}
		visited = append(visited, section...)
		area := len(section)
		sides := 0
		for _, index := range section {
			current := nodes[index]
			for _, delta := range corners {
				row, col := current.row+delta[0], current.col+delta[1]
				edge1 := outOfBounds(max_rows, max_cols, row, col) || slices.Index(section, coordToIndex(max_cols, row, col)) == -1
				row, col = current.row+delta[2], current.col+delta[3]
				corner := outOfBounds(max_rows, max_cols, row, col) || slices.Index(section, coordToIndex(max_cols, row, col)) == -1
				row, col = current.row+delta[4], current.col+delta[5]
				edge2 := outOfBounds(max_rows, max_cols, row, col) || slices.Index(section, coordToIndex(max_cols, row, col)) == -1
				//            . . .                         o . .                         . . o
				//            o X .                         o . .                         o X .
				//            o o .                         X o o                         o o .
				if (edge1 && corner && edge2) || (!edge1 && corner && !edge2) || (edge1 && !corner && edge2) {
					sides += 1
				}
			}
		}
		result += area * sides
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
