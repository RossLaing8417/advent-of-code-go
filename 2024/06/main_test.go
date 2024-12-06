package main

import "testing"

type test_data struct {
	name     string
	input    string
	expected int
}

func testPart(t *testing.T, tests []test_data, fn func(data string) (int, error)) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := fn(test.input)
			if err != nil {
				t.Error(err)
			}
			if result != test.expected {
				t.Errorf("Expected %d, got: %d", test.expected, result)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	testPart(t, []test_data{
		{
			name: "Example",
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expected: 41,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data{
		{
			name: "Example",
			input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			expected: 6,
		},
	}, solvePart2)
}
