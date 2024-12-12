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
			name: "Example 1",
			input: `AAAA
BBCD
BBCC
EEEC`,
			expected: 140,
		},
		{
			name: "Example 2",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			expected: 772,
		},
		{
			name: "Example 3",
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			expected: 1930,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data{
		{
			name: "Example 1",
			input: `AAAA
BBCD
BBCC
EEEC`,
			expected: 80,
		},
		{
			name: "Example 2",
			input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`,
			expected: 436,
		},
		{
			name: "Example 3",
			input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`,
			expected: 236,
		},
		{
			name: "Example 4",
			input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`,
			expected: 368,
		},
		{
			name: "Example 5",
			input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`,
			expected: 1206,
		},
	}, solvePart2)
}
