package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
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
		if data, err = os.ReadFile("{{.Year}}/{{.Day}}/input.txt"); err != nil {
			return "", err
		}
	} else if data, err = os.ReadFile("input.txt"); err != nil {
		return "", err
	}
	return strings.Trim(strings.ReplaceAll(string(data), "\r", ""), "\n"), nil
}

func solvePart1(data string) (int, error) {
	return strconv.Atoi(data + "1")
}

func solvePart2(data string) (int, error) {
	return strconv.Atoi(data + "2")
}
