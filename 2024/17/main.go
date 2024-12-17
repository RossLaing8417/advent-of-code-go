package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

type opcode int

const (
	adv opcode = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

func main() {
	aoc.Execute("2024/17/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (string, error) {
	var A, B, C int
	var program string
	fmt.Sscanf(data, "Register A: %d\nRegister B: %d\nRegister C: %d\n\nProgram: %s", &A, &B, &C, &program)
	result := solve(program, A, B, C)
	return result, nil
}

func solvePart2(data string) (int, error) {
	var A, B, C int
	var program string
	fmt.Sscanf(data, "Register A: %d\nRegister B: %d\nRegister C: %d\n\nProgram: %s", &A, &B, &C, &program)
	result := 0
	queue := [][2]int{{1, 0}}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		for A = item[1]; A < item[1]+8; A += 1 {
			p := solve(program, A, B, C)
			if p == program[len(program)-((item[0]*2)-1):] {
				if p == program {
					result = A
					break
				}
				queue = append(queue, [2]int{item[0] + 1, A * 8})
			}
		}
		if result != 0 {
			break
		}
	}
	return result, nil
}

func solve(program string, A int, B int, C int) string {
	result := []int{}
	for i := 0; i < len(program); i += 4 {
		code := opcode(program[i] - '0')
		literal := int(program[i+2] - '0')
		combo := literal
		switch combo {
		case 4:
			combo = A
		case 5:
			combo = B
		case 6:
			combo = C
		}
		switch code {
		case adv:
			A = A / int(math.Pow(2, float64(combo)))
		case bxl:
			B ^= literal
		case bst:
			B = combo % 8
		case jnz:
			if A != 0 {
				i = (literal * 2) - 4
			}
		case bxc:
			B ^= C
		case out:
			result = append(result, combo%8)
		case bdv:
			B = A / int(math.Pow(2, float64(combo)))
		case cdv:
			C = A / int(math.Pow(2, float64(combo)))
		}
	}
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprintf("%v", result)), ","), "[]")
}
