// 274. H-Index
package h_index

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	for i, n := range citations {
		if n >= len(citations)-i {
			return len(citations) - i
		}
	}

	return 0
}
