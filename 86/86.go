// 86. Partition List
package partition_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	leftSentinel := &ListNode{}
	rightSentinel := &ListNode{}
	leftTail := leftSentinel
	rightTail := rightSentinel

	curr := head
	for curr != nil {
		// Link current node to the correct partition
		if curr.Val < x {
			leftTail.Next = curr
			leftTail = leftTail.Next
		} else {
			rightTail.Next = curr
			rightTail = rightTail.Next
		}
		curr = curr.Next
	}

	// Terminate the right partition
	rightTail.Next = nil
	// Link the two partitions
	leftTail.Next = rightSentinel.Next

	return leftSentinel.Next
}
