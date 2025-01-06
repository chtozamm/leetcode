package main

import (
	"container/heap"
	"math"
)

func pickGifts(gifts []int, k int) int64 {
	pq := MaxHeap(gifts)
	heap.Init(&pq)

	for k > 0 {
		num := heap.Pop(&pq).(int)
		num = int(math.Floor(math.Sqrt(float64(num))))
		heap.Push(&pq, num)
		k--
	}

	var ans int64
	for _, num := range pq {
		ans += int64(num)
	}

	return ans
}
