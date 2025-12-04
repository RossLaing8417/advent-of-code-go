package main

import (
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/04/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	bdata, rowlen := []byte(data), strings.IndexByte(data, '\n')
	collen := (len(bdata) + 1) / (rowlen + 1)
	for i := range bdata {
		if bdata[i] != '@' {
			continue
		}
		count := countNeighbours(bdata, collen, rowlen, i)
		if count < 4 {
			result += 1
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	bdata, rowlen := []byte(data), strings.IndexByte(data, '\n')
	collen := (len(bdata) + 1) / (rowlen + 1)
	for {
		delta := visit(bdata, collen, rowlen)
		if delta == 0 {
			break
		}
		result += delta
	}
	return result, nil
}

var rotate = []complex128{
	1, -1,
	1i, -1i,
	1 + 1i, 1 - 1i,
	-1 + 1i, -1 - 1i,
}

func countNeighbours(data []byte, collen, rowlen, index int) int {
	pos, _ := index2Coord(rowlen, index)
	count := 0
	for _, dir := range rotate {
		i, ok := coord2Index(collen, rowlen, pos+dir)
		if ok && data[i] == '@' {
			count += 1
		}
	}
	return count
}

func coord2Index(collen, rowlen int, pos complex128) (int, bool) {
	col := int(real(pos))
	row := int(imag(pos))
	if col < 0 || col >= rowlen || row < 0 || row >= collen {
		return -1, false
	}
	return (row * (rowlen + 1)) + col, true
}

func index2Coord(rowlen, index int) (complex128, bool) {
	row := index / (rowlen + 1)
	col := index % (rowlen + 1)
	if col >= rowlen {
		return -1, false
	}
	return complex(float64(col), float64(row)), true
}

func visit(data []byte, collen, rowlen int) int {
	delta := 0
	for i := range data {
		if data[i] != '@' {
			continue
		}
		count := countNeighbours(data, collen, rowlen, i)
		if count < 4 {
			data[i] = 'x'
			delta += 1
		}
	}
	return delta
}
