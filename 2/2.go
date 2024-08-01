// 2. Add Two Numbers
package two_nums

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	curr := dummyHead
	carry := 0

	// Iterate through both linked lists
	for l1 != nil || l2 != nil || carry > 0 {
		// Get the values from the current nodes, or 0 if the node is nil
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// Calculate the sum and the carry
		sum := val1 + val2 + carry
		carry = sum / 10
		newNode := &ListNode{Val: sum % 10}

		// Add the new node to the result
		curr.Next = newNode
		curr = curr.Next
	}

	// Return the next of dummyHead as it points to the start of the result list
	return dummyHead.Next
}
