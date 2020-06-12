package insertdelete

import "math/rand"

// We use a struct to hold the elements in a slice, in order to return a random one, and in a map to do a constant
// time lookup as well as deleting the correct element.
type RandomizedSet struct {
	elements     []int
	elementToIdx map[int]int
}

// T: O(1)
// S: O(1)
func Constructor() RandomizedSet {
	return RandomizedSet{
		elements:     []int{},
		elementToIdx: map[int]int{},
	}
}

// To insert, we simply check the map and if the element exists, we return false.
// Otherwise, we append the new element to the slice and add it to the map. It's index will of course be
// the last one.
//
// T: O(1)
// S: O(1)
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.elementToIdx[val]; ok {
		return false
	}
	this.elementToIdx[val] = len(this.elementToIdx)
	this.elements = append(this.elements, val)
	return true
}

// If a value doesn't exist, we return false. If a value exists, we need to do few things.
// There are some edge cases for when we delete the only element left or the last element.
// In the other cases, the only way of having a constant time deletion, is to swap the target element with the last one,
// update last one's index to be the one of the target and finally delete the target from the end of the slice.
//
// T: O(1)
// S: O(1)
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.elementToIdx[val]
	if !ok {
		return false
	}
	lastIdx := len(this.elements) - 1
	if lastIdx != idx {
		lastElement := this.elements[lastIdx]
		this.elements[idx] = this.elements[lastIdx]
		if len(this.elements) > 0 {
			this.elementToIdx[lastElement] = idx
		}
	}
	delete(this.elementToIdx, val)
	this.elements = this.elements[:lastIdx]
	return true
}

// We return a random element from the slice based on the length of it.
//
// T: O(1)
// S: O(1)
func (this *RandomizedSet) GetRandom() int {
	return this.elements[rand.Intn(len(this.elements))]
}
