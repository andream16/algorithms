package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	res := fizzBuzz(15)
	if len(expected) != len(res) {
		t.Fatalf("expected result lenght %d, got %d", len(expected), len(res))
	}
	for i := 0; i < len(res); i++ {
		if expected[i] != res[i] {
			t.Fatalf("expected element at position %d to be %s, got %s", i, expected[i], res[i])
		}
	}
}
