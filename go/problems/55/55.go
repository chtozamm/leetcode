// 55. Jump Game
package jump_game

func canJump(nums []int) bool {
	maxReach := nums[0]

	for i, n := range nums {
		if maxReach < i {
			return false
		}
		if maxReach < i+n {
			maxReach = i + n
		}
		// Early exit if we can reach the last index
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return true
}
