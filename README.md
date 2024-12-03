# advent-of-code-go

Solving [Advent of Code](https://adventofcode.com/) in go

## Usage

A make file has been provided for running, testing, and generating solution files. You can run `make help` for more info

### Generating solution files

Generating today's solution files
```sh
make gen
```

Generating a specific year and/or day solution files
```sh
make gen year=<year> day=<day>
```

### Testing a solution

Testing today's solution (both parts)
```sh
make test
```

Testing today's solution (specific part)
```sh
make test part=<part>
```

Testing a specific year and day
```sh
make test year=<year> day=<day>
```

Testing a specific year and day with a specific part
```sh
make test year=<year> day=<day> part=<part>
```

### Running a solution

Running today's solution
```sh
make run part=<part>
```

Running a specific year and day
```sh
make run year=<year> day=<day> part=<part>
```

### Note
- A `.aoc_session` file containing your session cookie can be provided in the root dir for auto downloading the input when generating the solution files
- Solution files will not be overwritten
- The year will default to last year if the month is not December
