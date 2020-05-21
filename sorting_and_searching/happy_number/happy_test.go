package happy

import "testing"

func TestIsHappy(t *testing.T) {
	t.Run("19 should be an happy number", func(t *testing.T) {
		if !isHappy(19) {
			t.Fatal("19 is an happy number")
		}
	})
	t.Run("21 should not be an happy number", func(t *testing.T) {
		if isHappy(21) {
			t.Fatal("21 is not an happy number")
		}
	})
}
