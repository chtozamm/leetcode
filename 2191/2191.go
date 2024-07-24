// 2191. Sort the Jumbled Numbers
package jumbled_numbers

import (
	"sort"
	"strconv"
	"strings"
)

func sortJumbled(mapping []int, nums []int) []int {
	type remappedNum struct {
		original int
		remapped int
	}

	remappedNums := []remappedNum{} // Slice to store original numbers along with their remapped values

	// Remap each number according to the mapping
	for _, n := range nums {
		nStr := strconv.Itoa(n)
		digits := strings.Split(nStr, "")
		sb := strings.Builder{}

		// Remap each digit according to the provided mapping
		for _, d := range digits {
			dInt, _ := strconv.Atoi(d)
			mappedDigit := mapping[dInt]
			mappedDigitStr := strconv.Itoa(mappedDigit)
			sb.WriteString(mappedDigitStr)
		}

		rn, _ := strconv.Atoi(sb.String())
		remappedNums = append(remappedNums, remappedNum{original: n, remapped: rn})
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
