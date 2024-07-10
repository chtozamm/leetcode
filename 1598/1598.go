// 1598. Crawler Log Folder
package crawler

import "strings"

func minOperations(logs []string) int {
	stack := make([]string, 0, len(logs))
	for _, log := range logs {
		switch {
		case strings.HasPrefix(log, "../"):
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case strings.HasPrefix(log, "./"):
			continue
		default:
			stack = append(stack, log)
		}
	}
	return len(stack)
}
