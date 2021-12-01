package main

import (
	"fmt"
	"sort"
)

func minMeetingRooms(intervals [][]int) int {
	start, end := make([]int, len(intervals)), make([]int, len(intervals))

	for i := range intervals {
		start[i] = intervals[i][0]
		end[i] = intervals[i][1]
	}
	sort.Ints(start)
	sort.Ints(end)
	startIndex, endIndex := 0, 0
	for startIndex < len(intervals) {
		if start[startIndex] >= end[endIndex] {
			endIndex++
		}
		startIndex++
	}
	return startIndex - endIndex
}

func main() {
	a := [][]int{{0, 30}, {5, 10}, {15, 20}}
	res := minMeetingRooms(a)
	fmt.Println(res)
}
