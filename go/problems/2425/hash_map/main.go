package main

func xorAllNums(nums1 []int, nums2 []int) int {
	n := len(nums1)
	m := len(nums2)

	counts := make(map[int]int)

	for _, n1 := range nums1 {
		counts[n1] = counts[n1] + m
	}

	for _, n2 := range nums2 {
		counts[n2] = counts[n2] + n
	}

	ans := 0

	for num, count := range counts {
		if count%2 == 1 {
			ans ^= num
		}
	}

	return ans
}
