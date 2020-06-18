package hash

import "container/list"

type Entry struct {
	Key int
	Val int
}

type MyHashMap struct {
	Cap     int
	Buckets []*list.List
}

// The bigger the cap, the more buckets we can have, the more efficient the lookup in them is.
// Not using prime numbers will result into a worst distribution because of more factors -> collisions.
// We use a linked list in the buckets to make a very quick lookup.
//
// T: O(1)
// S: O(maxCap)
func Constructor() MyHashMap {
	const maxCap = 7919
	return MyHashMap{Cap: maxCap, Buckets: make([]*list.List, maxCap)}
}

// We check if we have a bucket for the hashed key and:
// - if not, we create a new one and append to it a new Entry
// - if we do, we lookup for the passed key in the current bucket and update it
//
// T: O(bucket) -> because in the worst case we need to arrive at the end of the list.
// S: O(1)
func (this *MyHashMap) Put(key int, value int) {
	hashedKey := this.hash(key)
	buckets := this.Buckets[hashedKey]
	if nil == buckets {
		this.Buckets[hashedKey] = list.New()
	} else {
		for b := buckets.Front(); b != nil; b = b.Next() {
			entry := b.Value.(Entry)
			if key == entry.Key {
				entry.Val = value
				b.Value = entry
				return
			}
		}
	}
	this.Buckets[hashedKey].PushBack(Entry{Key: key, Val: value})
}

// We check if we have a bucket for the hashed key and:
// - if not, it means that the key does not exists
// - if we do, we lookup for the passed key in the current bucket and return its value if found
//
// T: O(bucket) -> because in the worst case we need to arrive at the end of the list.
// S: O(1)
func (this *MyHashMap) Get(key int) int {
	hashedKey := this.hash(key)
	buckets := this.Buckets[hashedKey]
	if nil == buckets {
		return -1
	}
	for b := buckets.Front(); b != nil; b = b.Next() {
		entry := b.Value.(Entry)
		if key == entry.Key {
			return entry.Val
		}
	}
	return -1
}

// We check if we have a bucket for the hashed key and:
// - if not, it means that the key does not exists and we can return
// - if we do, we lookup for the passed key in the current bucket and delete it if found
//
// T: O(bucket) -> because in the worst case we need to arrive at the end of the list.
// S: O(1)
func (this *MyHashMap) Remove(key int) {
	hashedKey := this.hash(key)
	buckets := this.Buckets[hashedKey]
	if nil == buckets {
		return
	}
	for b := buckets.Front(); b != nil; b = b.Next() {
		entry := b.Value.(Entry)
		if key == entry.Key {
			buckets.Remove(b)
			return
		}
	}
}

// hash just uses the modulo operator. This is better with big prime numbers.
func (this *MyHashMap) hash(key int) int {
	return key % this.Cap
}
