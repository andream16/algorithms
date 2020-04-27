class Solution:

    def word_squares(self, words):

        self.words = words
        self.N = len(words[0])
        self.build_trie(self.words)

        results = []
        word_squares = []
        for word in words:
            word_squares = [word]
            self.backtracking(1, word_squares, results)
        return results

    def build_trie(self, words):
        self.trie = {}

        for wordIndex, word in enumerate(words):
            node = self.trie
            for char in word:
                if char in node:
                    node = node[char]
                else:
                    new_node = {}
                    new_node['#'] = []
                    node[char] = new_node
                    node = new_node
                node['#'].append(wordIndex)

    def backtracking(self, step, word_squares, results):
        if step == self.N:
            results.append(word_squares[:])
            return

        prefix = ''.join([word[step] for word in word_squares])
        for candidate in self.get_words_with_prefix(prefix):
            word_squares.append(candidate)
            self.backtracking(step+1, word_squares, results)
            word_squares.pop()

    def get_words_with_prefix(self, prefix):
        node = self.trie
        for char in prefix:
            if char not in node:
                return []
            node = node[char]
        return [self.words[wordIndex] for wordIndex in node['#']]


def main():
  solution = Solution()
  print(solution.word_squares(["area", "lead", "wall", "lady", "ball"]))


if __name__ == '__main__':
    main()