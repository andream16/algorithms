package movezeroes

// TODO review
func moveZeroes(nums []int) {
	var (
		zeros = 0
		l     = len(nums)
	)
	for i := 0; i < len(nums); i++ {
		if l-zeros == i {
			return
		}
		if nums[i] == 0 {
			nums = append(append(nums[:i], nums[i+1:]...), 0)
			zeros++
			i--
		}
	}
}
