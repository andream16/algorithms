package minstack

import "math"

// We can solve this problem by maintaining two separated stacks. The first one is a classic stack, we keep appending
// elements to it as we go. The second one is a min stack. We append to it only if the current value is less than the
// minimum, that is the last element of the stack. To avoid a bug with duplicated elements, we push to the stack
// if the current value is less or equal than the previous one. We remove elements from the regular stack as usual but
// from the min stack only if the stack.Top() is equals to minStack.Top(). All operations are constant but the space
// grows linearly.
//
// T: O(1)
// S: O(n)
type MinStack struct {
	Elements    []int
	MinElements []int
}

// T: O(1)
func Constructor() MinStack {
	return MinStack{
		Elements:    []int{},
		MinElements: []int{math.MaxInt64},
	}
}

// T: O(1)
func (this *MinStack) Push(x int) {
	if this.MinElements[len(this.MinElements)-1] >= x {
		this.MinElements = append(this.MinElements, x)
	}
	this.Elements = append(this.Elements, x)
}

// T: O(1)
func (this *MinStack) Pop() {
	lastIdx, lastMinIdx := len(this.Elements)-1, len(this.MinElements)-1
	if this.Elements[lastIdx] == this.MinElements[lastMinIdx] {
		this.MinElements = this.MinElements[:lastMinIdx]
	}
	this.Elements = this.Elements[:lastIdx]
}

// T: O(1)
func (this *MinStack) Top() int {
	return this.Elements[len(this.Elements)-1]
}

// T: O(1)
func (this *MinStack) GetMin() int {
	return this.MinElements[len(this.MinElements)-1]
}
