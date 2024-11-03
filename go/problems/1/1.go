// 1. Two Sum
package two_sum

func twoSum(nums []int, target int) []int {
	indices := map[int]int{}

	// Populate the map of indices
	for i, n := range nums {
		indices[n] = i
	}

	// Find the complement
	for i, n := range nums {
		complement := target - n
		if j, ok := indices[complement]; ok && j != i {
			return []int{i, j}
		}
	}

	return nil
}
