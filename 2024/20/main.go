package main

import (
	"math"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/20/input.txt", solvePart1, solvePart2)
}

type coord struct {
	x int
	y int
}

var adjacent [4]coord = [4]coord{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
var THRESHOLD int = 100

func solvePart1(data string) (int, error) {
	return solve(data, 2), nil
}

func solvePart2(data string) (int, error) {
	return solve(data, 20), nil
}

func indexToCoord(max_cols int, index int) coord {
	y := index / (max_cols + 1)
	return coord{index - y - (y * max_cols), y}
}

func solve(data string, distance int) int {
	max_x := strings.Index(data, "\n")
	grid := make(map[coord]byte)
	visited := make(map[coord]bool)
	var start coord
	var end coord
	for i := 0; i < len(data); i += 1 {
		if data[i] == '\n' {
			continue
		}
		coord := indexToCoord(max_x, i)
		grid[coord] = data[i]
		visited[coord] = false
		if data[i] == 'S' {
			start = coord
		} else if data[i] == 'E' {
			end = coord
		}
	}
	queue := []coord{start}
	dist := make(map[coord]int)
	dist[start] = 0
	path := []coord{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			break
		}
		if visited[current] {
			continue
		}
		visited[current] = true
		for _, delta := range adjacent {
			next := coord{current.x + delta.x, current.y + delta.y}
			if visited[next] || grid[next] == '#' {
				continue
			}
			dist[next] = dist[current] + 1
			path = append(path, next)
			queue = append(queue, next)
		}
	}
	result := 0
	for i, key := range path[:len(path)-1] {
		value := dist[key]
		for _, k := range path[i+1:] {
			v := dist[k]
			if v <= value {
				continue
			}
			d := int(math.Abs(float64(k.x-key.x))) + int(math.Abs(float64(k.y-key.y)))
			if d <= distance && v-value-d >= THRESHOLD {
				result += 1
			}
		}
	}
	return result
}
