package mergesortedarrays

// As we know that len(nums1) == n+m, we can start from the last elements of each array and compare them.
// nums1[i+j+1] is nums1[i] is this is bigger than nums2[j], else, it's nums2[j].
// We decrease i and j, respectively, when needed.
//
// T: O(n+m)
// S: O(1)
func merge(nums1 []int, m int, nums2 []int, n int) {

	for i, j := m-1, n-1; j >= 0; {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
			continue
		}
		nums1[i+j+1] = nums2[j]
		j--
	}

}
