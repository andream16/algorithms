package subarraysum

// To understand this problem, we need to look it from another point of view.
// As we are looking for contiguous elements that summed up give k, we can store the sum at each point in a map.
// The key is going to be the sum at a given element and the value is going to be how many times we saw that sum.
// When we consider an entry, we increase the sum and then if the map contains an element given by sum-k, it means that
// if we remove that entry, we can have a sum that reaches k. This is true for all the possible duplicates.
//
// Very good explanation here https://www.youtube.com/watch?v=HbbYPQc-Oo4
//
// T: O(n)
// S: O(n)
func subarraySum(nums []int, k int) int {

	var (
		counter int
		sum     int
		sums    = make(map[int]int, len(nums))
	)

	sums[0] = 1

	for _, n := range nums {
		sum += n
		if c, ok := sums[sum-k]; ok {
			counter += c
		}
		sums[sum]++
	}

	return counter
}
