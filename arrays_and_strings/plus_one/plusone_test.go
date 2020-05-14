package plusone

import "testing"

func TestPlusOne(t *testing.T) {
	for _, testCase := range []struct {
		given    []int
		expected []int
	}{
		{given: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{given: []int{4, 3, 2, 1}, expected: []int{4, 3, 2, 2}},
		{given: []int{9, 9, 9}, expected: []int{1, 0, 0, 0}},
	} {
		res := plusOne(testCase.given)
		for i := 0; i < len(res); i++ {
			if testCase.expected[i] != res[i] {
				t.Fatalf("expected %v but got %v", testCase.expected, res)
			}
		}
	}
}
