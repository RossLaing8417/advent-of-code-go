package main

import (
	"fmt"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/13/input.txt", solvePart1, solvePart2)
}

type coord struct {
	x int
	y int
}

func solvePart1(data string) (int, error) {
	result := 0
	for _, game := range strings.Split(data, "\n\n") {
		a, b, p := coord{}, coord{}, coord{}
		lines := strings.Split(game, "\n")
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &a.x, &a.y)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &b.x, &b.y)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &p.x, &p.y)
		// A*a.x + B*b.x = p.x
		// A*a.y + B.b.y = p.y
		det := (a.x*b.y - a.y*b.x)
		A := (p.x*b.y - p.y*b.x) / det
		B := (a.x*p.y - a.y*p.x) / det
		if A*a.x+B*b.x == p.x && A*a.y+B*b.y == p.y {
			result += 3*A + B
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	for _, game := range strings.Split(data, "\n\n") {
		a, b, p := coord{}, coord{}, coord{}
		lines := strings.Split(game, "\n")
		fmt.Sscanf(lines[0], "Button A: X+%d, Y+%d", &a.x, &a.y)
		fmt.Sscanf(lines[1], "Button B: X+%d, Y+%d", &b.x, &b.y)
		fmt.Sscanf(lines[2], "Prize: X=%d, Y=%d", &p.x, &p.y)
		p.x, p.y = p.x+10000000000000, p.y+10000000000000
		// A*a.x + B*b.x = p.x
		// A*a.y + B.b.y = p.y
		det := (a.x*b.y - a.y*b.x)
		A := (p.x*b.y - p.y*b.x) / det
		B := (a.x*p.y - a.y*p.x) / det
		if A*a.x+B*b.x == p.x && A*a.y+B*b.y == p.y {
			result += 3*A + B
		}
	}
	return result, nil
}
