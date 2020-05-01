package main

import "fmt"

type SnapshotArray struct {
	snaps       int
	currentSnap map[int]int
	snapHistory []map[int]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snaps:       0,
		currentSnap: make(map[int]int, length),
		snapHistory: []map[int]int{},
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.currentSnap[index] = val
}

func (this *SnapshotArray) Snap() int {
	this.snapHistory = append(this.snapHistory, this.currentSnap)
	this.currentSnap = deepCopySnap(this.currentSnap)
	this.snaps++
	return this.snaps - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	if len(this.snapHistory) < snap_id-1 {
		return 0
	}
	snap := this.snapHistory[snap_id]
	return snap[index]
}

func deepCopySnap(snap map[int]int) map[int]int {
	snapCopy := make(map[int]int, len(snap))
	for k, v := range snap {
		snapCopy[k] = v
	}
	return snapCopy
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

func main() {
	snap := Constructor(1)
	snap.Set(0, 15)
	fmt.Println(snap.Snap())    // 0
	fmt.Println(snap.Snap())    // 1
	fmt.Println(snap.Snap())    // 2
	fmt.Println(snap.Get(0, 2)) // 15
	fmt.Println(snap.Snap())    // 3
	fmt.Println(snap.Snap())    // 4
	fmt.Println(snap.Get(0, 0)) // 15
}
