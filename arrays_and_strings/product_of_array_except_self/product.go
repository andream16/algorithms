package product

// This problem can be solved using a greedy approach by using a two pass approach.
// We create a result array with same length's as the input one.
// On the first pass, we fill each position after the first one (that will be 1 by default) with the product of the
// previous product and the previous element in the input.
// On the second pass, we finalise the product array by multiplying each product element for the multiplication of the
// elements on its right.
//
// E.g. given input = [1, 2, 3, 4], the result should be [24, 12, 8, 6]
//
// - 1st pass:
// 	- prod[0] = 1 -> [1, 0, 0, 0]
//	- prod[1] = prod[0] * input[0] -> 1 * 1 -> [1, 1, 0, 0]
//	- prod[2] = prod[1] * input[1] -> 1 * 2 -> [1, 1, 2, 0]
//	- prod[3] = prod[2] * input[2] -> 2 * 3 -> [1, 1, 2, 6]
//
// - 2nd pass: input = [1, 2, 3, 4], prod = [1, 1, 2, 6], right = 1
// 	- prod[3] = prod[3] * right -> 6 * 1 -> 6 -> [1, 1, 2, 6]; right *= input[3] -> right *= 4 -> 4
// 	- prod[2] = prod[2] * right -> 2 * 4 -> 8 -> [1, 1, 8, 6]; right *= input[2] -> right *= 3 -> 12
// 	- prod[1] = prod[1] * right -> 1 * 12 -> 12 -> [1, 12, 8, 6]; right *= input[1]-> right *= 2 -> 24
// 	- prod[0] = prod[0] * right -> 1 * 24 -> 24 -> [24, 12, 8, 6]; right *= input[0]-> right *= 1 -> 24
//
// T: O(n)
// S: O(n)
func productExceptSelf(nums []int) []int {

	var (
		l        = len(nums)
		right    = 1
		products = make([]int, l)
	)

	// We know that the first value has to be 1.
	products[0] = 1

	// Initialise with products of numbers on the left of the current one.
	for i := 1; i < len(nums); i++ {
		products[i] = nums[i-1] * products[i-1]
	}

	// Finalise the result by multiplying what's left on current number's right.
	for i := l - 1; i >= 0; i-- {
		products[i] = products[i] * right
		right *= nums[i]
	}

	return products
}
