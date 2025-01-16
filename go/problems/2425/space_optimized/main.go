package main

func xorAllNums(nums1 []int, nums2 []int) int {
	xor1 := 0
	xor2 := 0

	n := len(nums1)
	m := len(nums2)

	if m%2 == 1 {
		for _, num := range nums1 {
			xor1 ^= num
		}
	}

	if n%2 == 1 {
		for _, num := range nums2 {
			xor2 ^= num
		}
	}

	return xor1 ^ xor2
}
