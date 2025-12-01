package main

import (
	"testing"
)

type test_data[T comparable] struct {
	name     string
	input    string
	expected T
}

func testPart[T comparable](t *testing.T, tests []test_data[T], fn func(data string) (T, error)) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := fn(test.input)
			if err != nil {
				t.Error(err)
			}
			if result != test.expected {
				t.Errorf("Expected %v, got: %v", test.expected, result)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	testPart(t, []test_data[int]{
		{
			name: "Example",
			input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			expected: 3,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data[int]{
		{
			name: "Example",
			input: `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			expected: 6,
		},
	}, solvePart2)
}
