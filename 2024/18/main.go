package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/18/input.txt", solvePart1, solvePart2)
}

var MAX_X int = 70 + 1
var MAX_Y int = 70 + 1
var MAX_I int = 1024

type coord struct {
	x int
	y int
}

func solvePart1(data string) (int, error) {
	grid := make(map[coord]int, 0)
	dist := make(map[coord]int, 0)
	for y := 0; y < MAX_Y; y += 1 {
		for x := 0; x < MAX_X; x += 1 {
			grid[coord{x, y}] = 0
		}
	}
	for i, line := range strings.Split(data, "\n") {
		c := coord{}
		_, err := fmt.Sscanf(line, "%d,%d", &c.x, &c.y)
		if err != nil {
			return 0, err
		}
		grid[c] = -1
		if i >= MAX_I-1 {
			break
		}
	}
	result := solve(&grid, &dist)
	return result, nil
}

func solvePart2(data string) (string, error) {
	grid := make(map[coord]int, 0)
	dist := make(map[coord]int, 0)
	for y := 0; y < MAX_Y; y += 1 {
		for x := 0; x < MAX_X; x += 1 {
			grid[coord{x, y}] = 0
		}
	}
	result := "-1,-1"
	for i, line := range strings.Split(data, "\n") {
		c := coord{}
		_, err := fmt.Sscanf(line, "%d,%d", &c.x, &c.y)
		if err != nil {
			return "", err
		}
		grid[c] = -1
		if i > MAX_I-1 {
			for key := range dist {
				dist[key] = 0
			}
			for key := range grid {
				if grid[key] == 1 {
					grid[key] = 0
				}
			}
			if solve(&grid, &dist) == -1 {
				result = fmt.Sprintf("%d,%d", c.x, c.y)
				break
			}
		}
	}
	return result, nil
}

func distance(coord coord) int {
	return int(math.Sqrt(
		math.Pow(float64(MAX_X-1-coord.x), 2) + math.Pow(float64(MAX_Y-1-coord.y), 2),
	))
}

var adjacent [4]coord = [4]coord{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func solve(grid *map[coord]int, dist *map[coord]int) int {
	start := coord{0, 0}
	end := coord{MAX_X - 1, MAX_Y - 1}
	queue := []coord{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current == end {
			return (*dist)[current]
		}
		if (*grid)[current] == 1 {
			continue
		}
		(*grid)[current] = 1
		for _, delta := range adjacent {
			next := coord{current.x + delta.x, current.y + delta.y}
			if value, existing := (*grid)[next]; !existing || value == -1 {
				continue
			}
			(*dist)[next] = (*dist)[current] + 1
			queue = append(queue, next)
		}
	}
	return -1
}
