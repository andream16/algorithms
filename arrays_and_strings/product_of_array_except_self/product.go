package main

import "fmt"

// O(n) + O(n) -> O(n)
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

	fmt.Println(products)

	// Finalise the result by multiplying what's left on current number's right.
	for i := l - 1; i >= 0; i-- {
		products[i] = products[i] * right
		right *= nums[i]
	}

	return products
}

func getProduct(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	products := make([]int, len(nums))
	products[0] = 1

	for i:=1; i<len(nums); i++ {
		products[i] = nums[i-1] * products[i]
	}

	r := 1
	for i:=len(nums)-1; i > 0; i-- {
		products[i] = r * products[i]
		r *= nums[i]
	}

	return products

}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	// p[i] -> n-1 * p[i-1]
	// [1, 1, 2, 6]
	// r=24
	// [24, 12, 8, 6]
	// r = 1
	// p[i] = p[i] * r
	// r *= nums[i]
}

