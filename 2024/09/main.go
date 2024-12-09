package main

import (
	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2024/09/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	disk := []int{}
	for i := 0; i < len(data); i += 1 {
		value := int(data[i] - '0')
		for j := 0; j < value; j += 1 {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1)
			}
		}
	}
	for i := 0; i < len(disk); i += 1 {
		if disk[i] == -1 {
			disk[i] = disk[len(disk)-1]
			disk[len(disk)-1] = -1
			disk = trim(disk)
		}
	}
	for i, id := range disk {
		if id != -1 {
			result += i * id
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	disk := []int{}
	for i := 0; i < len(data); i += 1 {
		value := int(data[i] - '0')
		for j := 0; j < value; j += 1 {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1)
			}
		}
	}
	for i := len(disk) - 1; i > 0; i -= 1 {
		if disk[i] == -1 {
			continue
		}
		end := i + 1
		start := i
		for j := i - 1; j > 0; j -= 1 {
			if disk[j] != disk[i] {
				break
			}
			start = j
		}
		i = start
		for j := 0; j < start; j += 1 {
			if disk[j] != -1 {
				continue
			}
			empty_start := j
			for disk[j] == -1 {
				j += 1
			}
			if len(disk[empty_start:j]) >= len(disk[start:end]) {
				for k := range disk[start:end] {
					disk[empty_start+k] = disk[start+k]
					disk[start+k] = -1
				}
				break
			}
		}
	}
	for i, id := range disk {
		if id != -1 {
			result += i * id
		}
	}
	return result, nil
}

func trim(values []int) []int {
	for i := len(values) - 1; i >= 0; i -= 1 {
		if values[i] != -1 {
			return values[:i+1]
		}
	}
	return values
}
