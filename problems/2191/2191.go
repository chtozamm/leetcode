// 2191. Sort the Jumbled Numbers
package jumbled_numbers

import (
	"sort"
	"strconv"
	"strings"
)

func sortJumbled(mapping []int, nums []int) []int {
	// Define a struct to store original numbers and their remapped values for sorting purposes
	type remappedNum struct {
		original int
		remapped int
	}

	remappedNums := make([]remappedNum, 0, len(nums))

	for _, num := range nums {
		// Convert the original number to a string and split it into digits
		numStr := strconv.Itoa(num)
		digits := strings.Split(numStr, "")

		sb := strings.Builder{}

		for _, d := range digits {
			// Remap each digit according to the provided mapping and append it to a new string
			digit, _ := strconv.Atoi(d)
			mappedDigit := mapping[digit]
			mappedDigitStr := strconv.Itoa(mappedDigit)
			sb.WriteString(mappedDigitStr)
		}

		// Convert the remapped string back to an integer and store it along with the original number
		rn, _ := strconv.Atoi(sb.String())
		remappedNums = append(remappedNums, remappedNum{original: num, remapped: rn})
	}

	// Sort the remapped numbers based on their remapped value
	sort.SliceStable(remappedNums, func(i, j int) bool {
		return remappedNums[i].remapped < remappedNums[j].remapped
	})

	// Extract the original numbers in sorted order
	res := make([]int, 0, len(nums))
	for _, rn := range remappedNums {
		res = append(res, rn.original)
	}

	return res
}
