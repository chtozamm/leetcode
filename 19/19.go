// 19. Remove Nth Node From End of List
package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{Next: head}
	first := sentinel
	second := sentinel

	// Move first pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers until first reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the N-th node from the end
	second.Next = second.Next.Next

	return sentinel.Next
}
