package main

import "sort"

func findScore(nums []int) int64 {
	score := int64(0)
	marked := make(map[int]bool)

	type element struct {
		value int
		index int
	}

	elements := make([]element, len(nums))
	for i, num := range nums {
		elements[i] = element{num, i}
	}

	sort.Slice(elements, func(i, j int) bool {
		if elements[i].value == elements[j].value {
			return elements[i].index < elements[j].index
		}

		return elements[i].value < elements[j].value
	})

	for _, e := range elements {
		if !marked[e.index] {
			score += int64(e.value)
			marked[e.index] = true

			if e.index > 0 {
				marked[e.index-1] = true
			}

			if e.index < len(nums)-1 {
				marked[e.index+1] = true
			}
		}
	}

	return score
}
