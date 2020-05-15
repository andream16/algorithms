package strobogrammaticnumber

import "testing"

func TestIsStrobogrammatic(t *testing.T) {

	for _, testCase := range []struct{
		num string
		expected bool
	}{
		{"1", true},
		{"69", true},
		{"88", true},
		{"692", false},
	}{
		if res := isStrobogrammatic(testCase.num); res != testCase.expected {
			t.Fatalf("expected %v, got %v", testCase.expected, res)
		}
	}

}
