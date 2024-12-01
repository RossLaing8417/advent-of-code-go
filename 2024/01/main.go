package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	if err := execute(); err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}
}

func execute() error {
	part := flag.Uint("part", 0, "solution part")
	flag.Parse()
	if *part == 0 {
		return errors.New("'-part' is mandatory")
	} else if *part > 2 {
		return errors.New(fmt.Sprintf("'-part %d' is invalid, valid value are 1 or 2\n", *part))
	}

	data, err := readFile()
	if err != nil {
		return err
	}

	var result int
	if *part == 1 {
		result, err = solvePart1(data)
	} else {
		result, err = solvePart2(data)
	}
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func readFile() (string, error) {
	var data []byte
	if _, err := os.Stat("input.txt"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
		if data, err = os.ReadFile("2024/01/input.txt"); err != nil {
			return "", err
		}
	} else if data, err = os.ReadFile("input.txt"); err != nil {
		return "", err
	}
	return strings.Trim(strings.ReplaceAll(string(data), "\r", ""), "\n"), nil
}

func solvePart1(data string) (int, error) {
	left_nums := make([]int, strings.Count(data, "\n")+1)
	right_nums := make([]int, len(left_nums))
	for i, line := range strings.Split(data, "\n") {
		var (
			left  int
			right int
		)
		if _, err := fmt.Sscanf(line, "%d   %d", &left, &right); err != nil {
			return 0, err
		}
		println(left, right)
		left_nums[i] = left
		right_nums[i] = right
	}
	slices.Sort(left_nums)
	slices.Sort(right_nums)
	result := 0
	for i := range left_nums {
		result += int(math.Abs(float64(left_nums[i]) - float64(right_nums[i])))
	}
	return result, nil
}

func solvePart2(data string) (int, error) {
	left_nums := make([]int, strings.Count(data, "\n")+1)
	right_nums := make([]int, len(left_nums))
	for i, line := range strings.Split(data, "\n") {
		var (
			left  int
			right int
		)
		if _, err := fmt.Sscanf(line, "%d   %d", &left, &right); err != nil {
			return 0, err
		}
		println(left, right)
		left_nums[i] = left
		right_nums[i] = right
	}
	slices.Sort(left_nums)
	slices.Sort(right_nums)
	result := 0
	for _, left := range left_nums {
		for _, right := range right_nums {
			if left == right {
				result += left
			}
		}
	}
	return result, nil
}
