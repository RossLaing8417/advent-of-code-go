package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/14/input.txt", solvePart1, solvePart2)
}

var MAX_X int = 101
var MAX_Y int = 103

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func solvePart1(data string) (int, error) {
	quad := [2][2]int{}
	mid_x := MAX_X / 2
	mid_y := MAX_Y / 2
	for _, line := range strings.Split(data, "\n") {
		r := robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.dx, &r.dy)
		r.x += r.dx * 100
		r.y += r.dy * 100
		r.x %= MAX_X
		r.y %= MAX_Y
		if r.x < 0 {
			r.x += MAX_X
		}
		if r.y < 0 {
			r.y += MAX_Y
		}
		if r.x == mid_x || r.y == mid_y {
			continue
		}
		qy, qx := 0, 0
		if r.x > mid_x {
			qx = 1
		}
		if r.y > mid_y {
			qy = 1
		}
		quad[qy][qx] += 1
	}
	result := 1
	for _, q := range quad {
		result *= q[0]
		result *= q[1]
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	mid_x := MAX_X / 2
	mid_y := MAX_Y / 2
	robots := []robot{}
	for _, line := range strings.Split(data, "\n") {
		r := robot{}
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.x, &r.y, &r.dx, &r.dy)
		robots = append(robots, r)
	}
	result := 0
	min_f := math.MaxInt
	robots_copy := []robot{}
	mmax := MAX_X * MAX_Y
	for i := 0; i < mmax; i += 1 {
		quad := [2][2]int{}
		for r := range robots {
			robots[r].x += robots[r].dx
			robots[r].y += robots[r].dy
			robots[r].x %= MAX_X
			robots[r].y %= MAX_Y
			if robots[r].x < 0 {
				robots[r].x += MAX_X
			}
			if robots[r].y < 0 {
				robots[r].y += MAX_Y
			}
			if robots[r].x == mid_x || robots[r].y == mid_y {
				continue
			}
			qy, qx := 0, 0
			if robots[r].x > mid_x {
				qx = 1
			}
			if robots[r].y > mid_y {
				qy = 1
			}
			quad[qy][qx] += 1
		}
		factor := 1
		for _, q := range quad {
			factor *= q[0]
			factor *= q[1]
		}
		if factor < min_f {
			min_f = factor
			result = i + 1
			robots_copy = append(robots_copy[:0], robots...)
		}
	}
	// for y := 0; y < MAX_Y; y += 1 {
	// 	for x := 0; x < MAX_X; x += 1 {
	// 		if slices.IndexFunc(robots_copy, func(r robot) bool { return r.x == x && r.y == y }) == -1 {
	// 			fmt.Printf(" ")
	// 		} else {
	// 			fmt.Printf("+")
	// 		}
	// 	}
	// 	fmt.Printf("\n")
	// }
	return result, nil
}
