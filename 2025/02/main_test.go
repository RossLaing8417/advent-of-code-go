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
			name:     "Example",
			input:    `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
			expected: 1227775554,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	testPart(t, []test_data[int]{
		{
			name:     "Example",
			input:    `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`,
			expected: 4174379265,
		},
	}, solvePart2)
}
