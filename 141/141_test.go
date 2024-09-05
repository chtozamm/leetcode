package linked_list_cycle

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		head     *ListNode
		expected bool
	}{
		{
			head:     buildLinkedListWithCycle([]int{3, 2, 0, -4}, 1),
			expected: true,
		},
		{
			head:     buildLinkedListWithCycle([]int{1, 2}, 0),
			expected: true,
		},
		{
			head:     buildLinkedListWithCycle([]int{1}, -1),
			expected: false,
		},
	}

	for i, tc := range testCases {
		if got := hasCycle(tc.head); got != tc.expected {
			t.Errorf("Case %d: got %v, expected %v", i+1, got, tc.expected)
		}
	}
}

func buildLinkedListWithCycle(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{Val: values[0]}
	current := head
	var cycleNode *ListNode

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{Val: values[i]}
		current = current.Next
		if i == pos {
			cycleNode = current
		}
	}

	if pos == 0 {
		current.Next = head
	} else if cycleNode != nil {
		current.Next = cycleNode
	}

	return head
}
