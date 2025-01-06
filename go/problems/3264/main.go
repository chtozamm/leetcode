package main

import "container/heap"

func getFinalState(nums []int, k int, multiplier int) []int {
	pq := make(PriorityQueue, len(nums))
	for i, num := range nums {
		pq[i] = &Item{value: num, idx: i}
	}
	heap.Init(&pq)

	for k > 0 {
		item := heap.Pop(&pq).(*Item)
		item.value *= multiplier
		nums[item.idx] = item.value
		heap.Push(&pq, item)
		k--
	}

	return nums
}
