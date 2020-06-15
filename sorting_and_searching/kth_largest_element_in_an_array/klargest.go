package kthlargest

import "container/heap"

// This problem can be solved by using a min-heap.
// After building a min heap, we pop as many elements as k-1.
// We then return heap.Pop() to get the k-th one.
//
// T: O(log n)
// S: O(log n)
func findKthLargest(nums []int, k int) int {
	minHeap := MinHeap(nums)
	heap.Init(&minHeap)
	for minHeap.Len() > k {
		heap.Pop(&minHeap)
	}
	return heap.Pop(&minHeap).(int)
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// T: O(log n)
// S: O(log n)
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// T: O(log n) - because when we remove the root (min) we need to reorder.
// S: O(log n)
func (h *MinHeap) Pop() interface{} {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}
