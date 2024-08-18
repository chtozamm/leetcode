// 45. Jump Game II
package jump_game_ii

func jump(nums []int) int {
	var currentReach, maxReach, count int

	for i, n := range nums[:len(nums)-1] { // Loop until the second last index
		if maxReach < i+n {
			maxReach = i + n
		}

		if i == currentReach {
			count++
			currentReach = maxReach

			// Return total number of jumps if we can reach the last index
			if currentReach >= len(nums)-1 {
				return count
			}
		}
	}

	return 0
}
