package phonedirectory

type PhoneDirectory struct {
	maxLen int
	nums []int
}

// Numbers go from 0 to maxNumbers.
//
// T: O(n)
// S: O(n)

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
	nums := make([]int, maxNumbers)
	for i:=0; i<maxNumbers; i++ {
		nums[i] = i
	}
	return PhoneDirectory{nums: nums, maxLen: maxNumbers}
}

// We just get the first one and delete it from the slice.
//
// T: O(1)
// S: O(1)

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
	if len(this.nums) == 0 {
		return -1
	}
	num := this.nums[0]
	this.nums = this.nums[1:]
	return num
}

// Using a slice, this is not constant. We could have used a map for this but would have been more memory intensive.
//
// T: O(n)
// S: O(1)

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	for _, n := range this.nums {
		if number == n {
			return true
		}
	}
	return false
}

// Same as for check. To avoid having duplicates. It depends on the constraints.
//
// T: O(n)
// S: O(1)

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int)  {
	if len(this.nums) == this.maxLen || this.Check(number) {
		return
	}
	this.nums = append(this.nums, number)
}
