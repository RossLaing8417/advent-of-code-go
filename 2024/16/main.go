package main

import (
	"math"
	"slices"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/16/input.txt", solvePart1, solvePart2)
}

type node struct {
	index int
	dir   int
	score int
	path  []int
}

type key struct {
	index int
	dir   int
}

var neighbours [][]int = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func solvePart1(data string) (int, error) {
	max_cols := strings.Index(data, "\n")
	start := strings.Index(data, "S")
	end := strings.Index(data, "E")
	grid := []byte(data)
	queue := []node{{start, dirToIndex('>'), 0, nil}}
	visited := make(map[key]node)
	result := 0
	for len(queue) > 0 {
		slices.SortFunc(queue, func(a node, b node) int { return a.score - b.score })
		current := queue[0]
		queue = queue[1:]
		if current.index == end {
			result = current.score
			break
		}
		current_key := key{current.index, current.dir}
		if _, existing := visited[current_key]; existing {
			continue
		}
		visited[current_key] = current
		row, col := indexToCoord(max_cols, current.index)
		idx := coordToIndex(max_cols, row+neighbours[current.dir][0], col+neighbours[current.dir][1])
		if grid[idx] != '#' {
			queue = append(queue, node{idx, current.dir, current.score + 1, nil})
		}
		queue = append(queue,
			node{current.index, (current.dir + 1) % 4, current.score + 1000, nil},
			node{current.index, (current.dir + 3) % 4, current.score + 1000, nil})
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	max_cols := strings.Index(data, "\n")
	start := strings.Index(data, "S")
	end := strings.Index(data, "E")
	grid := []byte(data)
	queue := []node{{start, dirToIndex('>'), 0, []int{start}}}
	visited := make(map[key]int)
	paths := []node{}
	min_score := math.MaxInt
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.score > min_score {
			continue
		}
		key := key{current.index, current.dir}
		if score, existing := visited[key]; existing && score < current.score {
			continue
		}
		visited[key] = current.score
		if current.index == end {
			min_score = min(min_score, current.score)
			paths = append(paths, current)
			continue
		}
		row, col := indexToCoord(max_cols, current.index)
		idx := coordToIndex(max_cols, row+neighbours[current.dir][0], col+neighbours[current.dir][1])
		if grid[idx] != '#' {
			path := make([]int, len(current.path))
			copy(path, current.path)
			queue = append(queue, node{idx, current.dir, current.score + 1, append(path, idx)})
		}
		queue = append(queue,
			node{current.index, (current.dir + 1) % 4, current.score + 1000, current.path},
			node{current.index, (current.dir + 3) % 4, current.score + 1000, current.path})
	}
	unique := make(map[int]bool)
	for _, path_node := range paths {
		if path_node.score != min_score {
			continue
		}
		for _, path := range path_node.path {
			unique[path] = true
		}
	}
	return len(unique), nil
}

func indexToCoord(max_cols int, index int) (int, int) {
	row := index / (max_cols + 1)
	return row, index - row - (row * max_cols)
}

func coordToIndex(max_cols int, row int, col int) int {
	return (row * max_cols) + row + col
}

func indexToDir(index int) byte {
	switch index {
	case 0:
		return '^'
	case 1:
		return '<'
	case 2:
		return 'v'
	case 3:
		return '>'
	}
	return 0
}

func dirToIndex(dir byte) int {
	switch dir {
	case '^':
		return 0
	case '<':
		return 1
	case 'v':
		return 2
	case '>':
		return 3
	}
	return -1
}
