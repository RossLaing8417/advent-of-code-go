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
	MAX_X = 11
	MAX_Y = 7
	testPart(t, []test_data{
		{
			name: "Example",
			input: `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
			expected: 12,
		},
	}, solvePart1)
}
