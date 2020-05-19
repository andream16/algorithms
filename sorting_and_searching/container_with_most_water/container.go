package container

// We calculate the area between each line with the formulae min(left, right) * (right-left)
// We save the max every time. We move left if left < right and right if left >= right.
//
// T: O(n)
// S: O(1)
func maxArea(height []int) int {
	area, left, right := 0, 0, len(height)-1

	for left < right {
		area = max(area, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
			continue
		}
		right--
	}

	return area
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
