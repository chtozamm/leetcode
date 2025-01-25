package main

import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)

	sortedNums := make([]int, n)
	copy(sortedNums, nums)
	slices.Sort(sortedNums)

	// Group numbers based on the limit
	grouped := make([][]int, 0)
	currentGroup := []int{sortedNums[0]}

	for i := 1; i < n; i++ {
		if abs(sortedNums[i]-sortedNums[i-1]) > limit {
			grouped = append(grouped, currentGroup)
			currentGroup = []int{sortedNums[i]}
		} else {
			currentGroup = append(currentGroup, sortedNums[i])
		}
	}
	grouped = append(grouped, currentGroup)

	// Create a map to track the smallest available number in each group
	groupIndex := 0
	groupMap := make(map[int]int)

	for _, group := range grouped {
		for _, num := range group {
			groupMap[num] = groupIndex
		}
		groupIndex++
	}

	result := make([]int, n)
	for i, num := range nums {
		group := groupMap[num]
		result[i] = grouped[group][0] // Replace with the smallest number in the group
		grouped[group] = grouped[group][1:]
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
