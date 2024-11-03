// 1636. Sort Array by Increasing Frequency
package frequency_sort

import (
	"sort"
)

func frequencySort(nums []int) []int {
	frequencyMap := make(map[int]int)
	for _, n := range nums {
		frequencyMap[n]++
	}

	res := make([]int, len(nums))
	copy(res, nums)

	sort.Slice(res, func(a, b int) bool {
		if frequencyMap[res[a]] == frequencyMap[res[b]] {
			return res[a] > res[b]
		} else {
			return frequencyMap[res[a]] < frequencyMap[res[b]]
		}
	})

	return res
}
