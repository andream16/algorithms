package happy

// A number is happy if we find 1. We check if we are in a cycle by using an hashmap.
// To get the next number, we do, for n = 19:
// - d = 19 % 10 = 9; n = 19 / 10 = 1; next += d * d = 81
// - d = 1 % 10 = 1; n = 1 / 10 = 0; next += d * d = 82
//
// This implementation uses Floyd's Cycle-Finding Algorithm that uses two concepts, turtle and hare.
// The hare goes twice as fast as the turtle. If they meet it means that we have a cycle.
//
// T: O(log n)
// S: O(1)
func isHappy(n int) bool {

	slowRunner, fastRunner := n, getNext(n)

	for 1 != fastRunner && fastRunner != slowRunner {
		slowRunner, fastRunner = getNext(slowRunner), getNext(getNext(fastRunner))
	}

	return fastRunner == 1

}

func getNext(n int) int {
	next := 0
	for n > 0 {
		d := n % 10
		n /= 10
		next += d * d
	}
	return next
}

// A number is happy if we find 1. We check if we are in a cycle by using an hashmap.
// To get the next number, we do, for n = 19:
// - d = 19 % 10 = 9; n = 19 / 10 = 1; next += d * d = 81
// - d = 1 % 10 = 1; n = 1 / 10 = 0; next += d * d = 82
//
// T: O(log n)
// S: O(log n)
func isHappyBasic(n int) bool {
	var (
		seen = map[int]struct{}{}
		next int
	)

	for 1 != n {
		next = getNext(n)
		_, ok := seen[next]
		if ok {
			return false
		}
		seen[next] = struct{}{}
		n = next
	}

	return true
}
