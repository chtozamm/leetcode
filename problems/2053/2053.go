// 2053. Kth Distinct String in an Array
package distinct_string

func kthDistinct(arr []string, k int) string {
	counts := make(map[string]int)
	for _, s := range arr {
		counts[s]++
	}
	for _, s := range arr {
		if counts[s] == 1 {
			k--
			if k == 0 {
				return s
			}
		}
	}
	return ""
}
