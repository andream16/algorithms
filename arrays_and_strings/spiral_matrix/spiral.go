package spiral

// As we are interested in doing right > down > left > up and excluding already visited elements,
// we can simply use 4 counters, one for each direction. We go:
//  - right from first column to last available
//  - down from the next row to the last available
//  - left from the previous column to the leftmost available
//  - up from the previous row to the uppermost available
//
// Visual explanation https://leetcode.com/problems/spiral-matrix/Figures/54_spiralmatrix.png
//
// T: O(n)
// S: O(n)
func spiralOrder(matrix [][]int) []int {

	if 0 == len(matrix) {
		return make([]int, 0)
	}

	var (
		res            []int
		r1, r2, c1, c2 = 0, len(matrix) - 1, 0, len(matrix[0]) - 1
	)

	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			res = append(res, matrix[r1][c])
		}
		for r := r1 + 1; r <= r2; r++ {
			res = append(res, matrix[r][c2])
		}
		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				res = append(res, matrix[r2][c])
			}
			for r := r2; r > r1; r-- {
				res = append(res, matrix[r][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}

	return res
}
