package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/RossLaing8417/advent-of-code-go/aoc"
)

func main() {
	aoc.Execute("2025/05/input.txt", solvePart1, solvePart2)
}

func solvePart1(data string) (int, error) {
	result := 0
	r_start, r_end, ranges := []uint64{}, []uint64{}, true
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			ranges = false
			continue
		}
		if ranges {
			var start, end uint64
			_, err := fmt.Sscanf(line, "%d-%d", &start, &end)
			if err != nil {
				return 0, nil
			}
			r_start = append(r_start, start)
			r_end = append(r_end, end)
		} else {
			id, err := strconv.ParseUint(line, 10, 0)
			if err != nil {
				return 0, err
			}
			for i := range r_start {
				start, end := r_start[i], r_end[i]
				if id >= start && id <= end {
					result += 1
					break
				}
			}
		}
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	result := 0
	r_start, r_end, ranges := []uint64{}, []uint64{}, true
	for _, line := range strings.Split(data, "\n") {
		if line == "" {
			ranges = false
			continue
		}
		if ranges {
			var start, end uint64
			_, err := fmt.Sscanf(line, "%d-%d", &start, &end)
			if err != nil {
				return 0, nil
			}
			r_start = append(r_start, start)
			r_end = append(r_end, end)
		}
	}
	changed := true
	for changed {
		changed = false
		for i := range r_start {
			start_i, end_i := r_start[i], r_end[i]
			for j := range r_start {
				if j == i {
					continue
				}
				start_j, end_j := r_start[j], r_end[j]
				if start_j >= start_i && start_j <= end_i {
					changed = true
					if end_j > end_i {
						r_end[i] = end_j
					}
				}
				if end_j >= start_i && end_j <= end_i {
					changed = true
					if start_j < start_i {
						r_start[i] = start_j
					}
				}
				if changed {
					r_start = slices.Delete(r_start, j, j+1)
					r_end = slices.Delete(r_end, j, j+1)
					break
				}
			}
			if changed {
				break
			}
		}
	}
	for i := range r_start {
		start, end := r_start[i], r_end[i]
		result += int((end - start) + 1)
	}
	return result, nil
}
