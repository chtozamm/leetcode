// 350. Intersection of Two Arrays II
package array

func intersect(nums1 []int, nums2 []int) []int {
	frequencies := make(map[int]int)
	res := []int{}

	// Populate maps of number frequencies
	for _, num := range nums1 {
		frequencies[num]++
	}

	for _, num := range nums2 {
		if frequencies[num] > 0 {
			frequencies[num]--
			res = append(res, num)
		}
	}

	return res
}
