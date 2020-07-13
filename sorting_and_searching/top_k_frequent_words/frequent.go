package frequent

import "sort"

type freqWord struct {
	word      string
	frequency int
}

// To solve this problem, we rely on sorting the words by their frequency.
// If two words have the same frequency, we sort them lexicographically.
// To do so, we use freqWord that holds (word, frequency), a slice to hold such freqWord entries
// and a dictionary with word as key and index of the word in the slice as value. This allow us to quickly check
// if a word was visited already and update it's frequency in constant time.
// We fill the slice and the map in a greedy manner. Once this is done, we sort the slice by frequency and lexicographic order
// in case two entries' frequency is equal. Finally, we filter the top k words.
//
// T: O(n log n)
// S: O(n)
func topKFrequent(words []string, k int) []string {
	var (
		uniqueWords []string
		wordCount = map[string]int{}
	)
	for _, word := range words {
		if _, ok := wordCount[word]; !ok {
			uniqueWords = append(uniqueWords, word)
		}
		wordCount[word]++
	}

	sort.Slice(uniqueWords, func(i, j int) bool {
		if wordCount[uniqueWords[i]] == wordCount[uniqueWords[j]] {
			return uniqueWords[i] < uniqueWords[j]
		}
		return wordCount[uniqueWords[i]] > wordCount[uniqueWords[j]]
	})

	return uniqueWords[:k]
}
