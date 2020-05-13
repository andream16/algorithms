package single

// Using binary search we can find the number with O(log n).
//
// T: O(log n)
// S: O(1)
func singleNonDuplicate(nums []int) int {
	return binarySearch(0, len(nums), nums)
}

func binarySearch(left, right int, arr []int) int {
	if left == right {
		return 0
	}
	middle := (left + right) / 2
	if isSingle(middle, arr) {
		return arr[middle]
	}
	return binarySearch(left, middle, arr) + binarySearch(middle+1, right, arr)
}

func isSingle(index int, arr []int) bool {
	return (len(arr) == 1) ||
		(index > 0 && index < len(arr)-1 && arr[index] > arr[index-1] && arr[index] < arr[index+1]) ||
		(index == 0 && len(arr) > 1 && arr[index] < arr[index+1]) ||
		(index == len(arr)-1 && len(arr) > 1 && arr[index] > arr[index-1])
}
