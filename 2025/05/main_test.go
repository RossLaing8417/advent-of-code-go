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
	testPart(t, []test_data[int]{
		{
			name: "Example",
			input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			expected: 3,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data[int]{
		{
			name: "Example",
			input: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			expected: 14,
		},
	}, solvePart2)
}
