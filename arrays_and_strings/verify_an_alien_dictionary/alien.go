package alien

// We know that the words are sorted if w[i] <= w[i+1].
// To check of the order is valid, we need to compare w[i] with w[i+1] character by character.
// If c1 comes before c2 or w[i] and w[i+1] have the same prefix but w[i] is longer, then, the order is not valid.
//
// T: O((n-1) * |shortestWord|)
// S: O(order)
func isAlienSorted(words []string, order string) bool {

	orderIdx := make(map[rune]int, 26)
	for i, r := range order {
		orderIdx[r] = i
	}

	for i := 1; i < len(words); i++ {

		diff, word1, word2 := false, words[i-1], words[i]
		l1, l2 := len(word1), len(word2)

		for k := 0; k < min(l1, l2); k++ {
			c1, c2 := word1[k], word2[k]
			if c1 != c2 {
				diff = true
				if orderIdx[rune(c1)] > orderIdx[rune(c2)] {
					return false
				}
				break
			}
		}
		if !diff && l1 > l2 {
			return false
		}
	}

	return true
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
