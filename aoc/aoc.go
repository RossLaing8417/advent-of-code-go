package aoc

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func Execute(filepath string, part1_fn func(data string) (int, error), part2_fn func(data string) (int, error)) {
	if err := execute(filepath, part1_fn, part2_fn); err != nil {
		os.Stderr.Write([]byte(err.Error()))
		os.Exit(1)
	}
}

func execute(filepath string, part1_fn func(data string) (int, error), part2_fn func(data string) (int, error)) error {
	part := flag.Uint("part", 0, "solution part")
	flag.Parse()

	if *part > 2 {
		return errors.New(fmt.Sprintf("'-part %d' is invalid, valid value are 1 or 2\n", *part))
	}

	data, err := readFile(filepath)
	if err != nil {
		return err
	}

	if *part == 0 || *part == 1 {
		result, err := part1_fn(data)
		if err != nil {
			return err
		}
		if *part == 0 {
			fmt.Print("Part 1: ")
		}
		fmt.Print(result)
		if *part == 0 {
			fmt.Print("\n")
		}
	}
	if *part == 0 || *part == 2 {
		result, err := part2_fn(data)
		if err != nil {
			return err
		}
		if *part == 0 {
			fmt.Print("Part 2: ")
		}
		fmt.Print(result)
	}

	return nil
}

func readFile(filepath string) (string, error) {
	var data []byte
	if _, err := os.Stat("input.txt"); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return "", err
		}
		if data, err = os.ReadFile(filepath); err != nil {
			return "", err
		}
	} else if data, err = os.ReadFile("input.txt"); err != nil {
		return "", err
	}
	return strings.Trim(strings.ReplaceAll(string(data), "\r", ""), "\n"), nil
}
