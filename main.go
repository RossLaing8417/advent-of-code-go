package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func main() {
	if err := execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func execute() error {
	generate := flag.Bool("generate", false, "generate solution files")
	test := flag.Bool("test", false, "test solution")
	run := flag.Bool("run", false, "run solution")
	year := flag.Uint("year", 0, "solution year")
	day := flag.Uint("day", 0, "solution day")
	part := flag.Uint("part", 0, "solution part")
	flag.Parse()

	if (*generate && *test) || (*generate && *run) || (*test && *run) {
		return errors.New("'-generate', '-test,' and '-run' are mutually exclusive")
	}

	today := time.Now()

	if *year == 0 {
		*year = uint(today.Year())
		if today.Month() < 12 {
			*year -= 1
		}
	}

	if *day == 0 {
		*day = uint(today.Day())
	}

	if *year < 2015 {
		return errors.New("Year cannot be before 2015")
	} else if *year > uint(today.Year()) {
		return errors.New(fmt.Sprintln("Year cannot be after", today.Year()))
	}
	if *day > 25 {
		return errors.New("Day cannot be after 25")
	}

	if *part == 0 {
		if *run {
			return errors.New("'-part' is mandatory when '-run' is specified")
		}
	} else if *part > 2 {
		return errors.New(fmt.Sprintf("'-part %d' is invalid, valid value are 1 or 2", *part))
	}

	if *generate {
		if err := generateSolutionFiles(*year, *day); err != nil {
			return err
		}
	} else {
		dir := fmt.Sprintf("%4d/%02d", *year, *day)
		var cmd *exec.Cmd
		if *test {
			if *part == 0 {
				cmd = exec.Command(
					"go",
					"test",
					dir+"/main.go",
					dir+"/main_test.go",
				)
			} else {
				cmd = exec.Command(
					"go",
					"test",
					dir+"/main.go",
					dir+"/main_test.go",
					"-run",
					fmt.Sprintf("TestPart%d", *part),
				)
			}
		}
		if *run {
			cmd = exec.Command(
				"go",
				"run",
				dir+"/main.go",
				"-part",
				strconv.Itoa(int(*part)),
			)
		}
		output, err := cmd.Output()
		if err != nil {
			if exit_err, ok := err.(*exec.ExitError); ok && len(exit_err.Stderr) > 0 {
				fmt.Println(string(exit_err.Stderr))
			}
			if len(output) > 0 {
				fmt.Println(string(output))
			}
			return err
		}
		if len(output) > 0 {
			fmt.Println(string(output))
		}
	}

	return nil
}

func generateSolutionFiles(year uint, day uint) error {
	dir := fmt.Sprintf("%4d/%02d", year, day)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles("tmpl/main.go", "tmpl/main_test.go")
	if err != nil {
		return err
	}

	data := struct {
		Year string
		Day  string
	}{
		Year: fmt.Sprintf("%4d", year),
		Day:  fmt.Sprintf("%02d", day),
	}

	main_file := dir + "/main.go"
	if _, err := os.Stat(main_file); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}

		file, err := os.Create(main_file)
		if err != nil {
			return nil
		}

		if err := tmpl.ExecuteTemplate(file, "main.go", data); err != nil {
			return err

		}
	} else {
		return errors.New(fmt.Sprintf("File '%s' already exists", main_file))
	}

	main_test_file := dir + "/main_test.go"
	if _, err := os.Stat(main_test_file); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}

		file, err := os.Create(main_test_file)
		if err != nil {
			return nil
		}

		if err := tmpl.ExecuteTemplate(file, "main_test.go", data); err != nil {
			return err

		}
	} else {
		return errors.New(fmt.Sprintf("File '%s' already exists and will not be overwritten...", main_test_file))
	}

	session_data, err := os.ReadFile(".aoc_session")
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	if len(session_data) > 0 {
		context, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		request, err := http.NewRequestWithContext(context, "GET", fmt.Sprintf("https://adventofcode.com/%4d/day/%d/input", year, day), nil)
		if err != nil {
			return err
		}
		request.AddCookie(&http.Cookie{
			Name:  "session",
			Value: strings.Trim(strings.ReplaceAll(string(session_data), "\r", ""), "\n"),
		})
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return err
		}
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		os.WriteFile(dir+"/input.txt", body, 0644)
	} else {
		return errors.New(".aoc_session not found, skipping input download...")
	}

	return nil
}
