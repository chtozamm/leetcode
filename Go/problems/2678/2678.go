// 2678. Number of Senior Citizens
package seniors

import "strconv"

func countSeniors(details []string) int {
	count := 0
	for _, d := range details {
		ageStr := d[11:13]
		age, _ := strconv.Atoi(ageStr)
		if age > 60 {
			count++
		}
	}
	return count
}
