// 1395. Count Number of Teams
package count_teams

func numTeams(rating []int) int {
	n := len(rating)
	teams := 0

	// Iterate through each soldier as the middle soldier
	for mid := 0; mid < n; mid++ {
		leftSmaller := 0
		rightLarger := 0

		// Count soldiers with smaller ratings on the left side of the current soldier
		for left := mid - 1; left >= 0; left-- {
			if rating[left] < rating[mid] {
				leftSmaller++
			}
		}

		// Count soldiers with larger ratings on the right side of the current soldier
		for right := mid + 1; right < n; right++ {
			if rating[right] > rating[mid] {
				rightLarger++
			}
		}

		// Calculate and add the number of ascending rating teams (small-mid-large)
		teams += leftSmaller * rightLarger

		// Calculate soldiers with larger ratings on the left and smaller ratings on the right
		leftLarger := mid - leftSmaller
		rightSmaller := n - mid - 1 - rightLarger

		// Calculate and add the number of descending rating teams (large-mid-small)
		teams += leftLarger * rightSmaller
	}

	// Return the total number of valid teams
	return teams
}
