package main

import "testing"

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
	testPart(t, []test_data[string]{
		{
			name: "Example",
			input: `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`,
			expected: "4,6,3,5,6,3,5,2,1,0",
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data[int]{
		{
			name: "Example",
			input: `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`,
			expected: 117440,
		},
	}, solvePart2)
}
