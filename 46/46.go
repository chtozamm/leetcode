// 46. Permutations
package permutations

func permute(nums []int) [][]int {
	var permutations [][]int
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(nums) {
			perm := make([]int, len(nums))
			copy(perm, nums)
			permutations = append(permutations, perm)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	backtrack(0)
	return permutations
}
