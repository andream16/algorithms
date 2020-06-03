package meeting

import "sort"

// This problem can be solved in a naive way. We solve it like we would approach the problem in real life.
// We separate start and end times in their own arrays and then we order them. We start processing start times and
// compare them with the current end time. If the current start time is after the current end time, it means that one or
// more meetings have ended and we can re-use one available room. In the other case, we need another room.
//
// T: O(n log n)
// O: O(n)
func minMeetingRooms(intervals [][]int) int {
	var (
		numRooms   int
		startTimes = make([]int, len(intervals))
		endTimes   = make([]int, len(intervals))
	)

	for i := range intervals {
		startTimes[i] = intervals[i][0]
		endTimes[i] = intervals[i][1]
	}

	sort.Ints(startTimes)
	sort.Ints(endTimes)

	var startIdx, endIdx int

	for ; startIdx < len(startTimes); startIdx++ {
		if startTimes[startIdx] >= endTimes[endIdx] {
			endIdx++
			continue
		}
		numRooms++
	}

	return numRooms
}
