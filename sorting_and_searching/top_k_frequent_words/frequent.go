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
		freqWords     = []freqWord{}
		currFreqWords = map[string]int{}
		res           = make([]string, k)
	)

	for _, w := range words {
		idx, ok := currFreqWords[w]
		if ok {
			freqWords[idx].frequency++
			continue
		}
		freqWords = append(freqWords, freqWord{word: w, frequency: 1})
		currFreqWords[w] = len(freqWords) - 1
	}

	sort.Slice(freqWords, func(i, j int) bool {
		if freqWords[i].frequency == freqWords[j].frequency {
			return freqWords[i].word < freqWords[j].word
		}
		return freqWords[i].frequency > freqWords[j].frequency
	})

	for i := 0; i < k; i++ {
		res[i] = freqWords[i].word
	}

	return res

}
