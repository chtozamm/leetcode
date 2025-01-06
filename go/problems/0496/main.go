package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	hashmap := make(map[int]int)

	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < num {
			hashmap[stack[len(stack)-1]] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}

	ans := make([]int, 0, len(nums1))

	for _, num := range nums1 {
		if n, ok := hashmap[num]; ok {
			ans = append(ans, n)
		} else {
			ans = append(ans, -1)
		}
	}

	return ans
}
