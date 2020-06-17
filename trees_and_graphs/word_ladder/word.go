package wordladder

type leveledWord struct {
	val   string
	level int
}

type queue struct {
	values []*leveledWord
}

func (q *queue) enqueue(val string, level int) {
	q.values = append(q.values, &leveledWord{
		val:   val,
		level: level,
	})
}

func (q *queue) dequeue() *leveledWord {
	w := q.values[0]
	q.values = q.values[1:]
	return w
}

func (q *queue) empty() bool {
	return len(q.values) == 0
}

// T: O(N*L) for initial combinations where N is the number of words and L is the length of a word.
// O(N*L*W) for BFS, current combination and possibly checking the set associated.
// O(N*L) should be bigger.

// S: N*L + N or initial combinations and words set, N for the queue, N for visited.
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == "" || endWord == "" || len(wordList) == 0 {
		return 0
	}

	var (
		words        = make(map[string]struct{}, len(wordList))
		combinations = map[string]map[string]struct{}{}
		q            = &queue{}
		visited      = map[string]struct{}{}
	)

	// find all possible combinations by alternating "*"
	for _, word := range wordList {
		words[word] = struct{}{}
		for i := 0; i < len(word); i++ {
			combo := word[:i] + "*" + word[i+1:]
			combs, ok := combinations[combo]
			if ok {
				combs[word] = struct{}{}
				combinations[combo] = combs
				continue
			}
			combinations[combo] = map[string]struct{}{
				word: {},
			}
		}
	}

	// if endWord is not in the word list, we can return
	if _, ok := words[endWord]; !ok {
		return 0
	}

	// as the words are all linked together, we start with level 1 from begin word
	q.enqueue(beginWord, 1)
	// let's keep track of what we already visited to avoid infinite loops
	visited[beginWord] = struct{}{}

	for !q.empty() {
		// extract the current word
		word := q.dequeue()

		for i := 0; i < len(word.val); i++ {
			// check if the current combination exists in the combinations map
			combo := word.val[:i] + "*" + word.val[i+1:]
			for w := range combinations[combo] {
				// if so, if we find the endWord we are done
				if endWord == w {
					return word.level + 1
				}
				// otherwise we need to visit the next word, if it hasn't been visited yet
				if _, ok := visited[w]; !ok {
					visited[w] = struct{}{}
					q.enqueue(w, word.level+1)
				}
			}
			combinations[combo] = map[string]struct{}{}
		}

	}

	return 0
}
