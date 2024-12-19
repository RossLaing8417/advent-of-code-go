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
	MAX_X = 6 + 1
	MAX_Y = 6 + 1
	MAX_I = 12
	testPart(t, []test_data[int]{
		{
			name: "Example",
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
			expected: 22,
		},
	}, solvePart1)
}

func TestPart2(t *testing.T) {
	MAX_X = 6 + 1
	MAX_Y = 6 + 1
	MAX_I = 12
	testPart(t, []test_data[string]{
		{
			name: "Example",
			input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
			expected: "6,1",
		},
	}, solvePart2)
}
