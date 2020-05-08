package main

import (
	"fmt"
)

type MyCalendarTwo struct {
	Overlaps []Event
	Events   []Event
}

type Event struct {
	start int
	end   int
}

// T: O(1)
// S: O(1)
func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		Overlaps: []Event{},
		Events:   []Event{},
	}
}

// T: O(N^2) because for each event we visit every previous one.
// S: O(N) for events + O(L) for number of overlaps
func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, event := range this.Overlaps {
		if event.overlaps(start, end) {
			return false
		}
	}

	for _, event := range this.Events {
		if event.overlaps(start, end) {
			this.Overlaps = append(this.Overlaps, Event{start: max(event.start, start), end: min(event.end, end)})
		}
	}

	this.Events = append(this.Events, Event{start: start, end: end})
	return true
}

func (i Event) overlaps(start, end int) bool {
	// reverse for when they don't overlap
	return i.end > start && end > i.start
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	calendar := Constructor()
	fmt.Println(calendar.Book(10, 20)) // true
	fmt.Println(calendar.Book(50, 60)) // true
	fmt.Println(calendar.Book(10, 40)) // true
	fmt.Println(calendar.Book(5, 15))  // false
	fmt.Println(calendar.Book(5, 10))  // true
	fmt.Println(calendar.Book(25, 55)) // true
}
