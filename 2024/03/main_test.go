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
			name:     "Example",
			input:    `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			expected: 161,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data{
		{
			name:     "Example Part 1",
			input:    `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			expected: 161,
		},
		{
			name:     "Example Part 2",
			input:    `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			expected: 48,
		},
	}, solvePart2)
}
