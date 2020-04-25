package main

import "fmt"

func wordSquares(words []string) [][]string {
	var (
		m = make([][]rune, len(words))
		letterToWord = make(map[rune]string, len(words))
		wm = make(map[string]struct{}, len(words))
		squares = [][]string{}
	)

	for i:= 0; i < len(words); i++ {
		row := make([]rune, len(words[i]))
		for j:=0; j<len(words[i]); j++ {
			row[j] = rune(words[i][j])
		}
		m[i] = row
		letterToWord[rune(words[i][0])] = words[i]
		wm[words[i]] = struct{}{}
	}

	originalM := m

	for i:=0; i<len(m); i++ {
		for j:=1; j<len(m); j++ {
			if isSquare(m, wm) {
				squares = append(squares, getWordsInMatrix(m))
			}
			m = append(m[:i], m[j], m[i:j][0], m[j+1:][0])
		}
		m = originalM
	}

	return squares

}

func isSquare(m [][]rune, wm map[string]struct{}) bool {
	for i:=0; i<len(m); i++ {
		colWord := ""
		for j:=0; j<len(m[i]); j++ {
			colWord += string(m[j][i])
		}
		if _, ok := wm[colWord]; !ok || colWord != string(m[i]) {
			return false
		}
	}
	return true
}

func getWordsInMatrix(m [][]rune) []string {
	words := make([]string, len(m))
	for i:=0; i<len(m); i++ {
		words[i] = string(m[i])
	}
	return words
}

// [[0], [1], [2], [3]]

func swapRows(r1, r2 int, m [][]rune) {
	// arr[:r1][0], arr[r2], arr[r1:r2][0], arr[r2+1:][0]
	m = append(m[:r1], m[r2], m[r1:r2][0], m[r2+1:][0])
}

// area
// lead
// wall
// lady

// ball

// Input:
// ["area","lead","wall","lady","ball"]
//
// Output:
// [
//   [ "wall",
//    "area",
//    "lead",
//    "lady"
//   ],
//   [ "ball",
//    "area",
//    "lead",
//    "lady"
//   ]
// ]
func main() {
	fmt.Println(wordSquares([]string{"area","lead","wall","lady"}))
}
