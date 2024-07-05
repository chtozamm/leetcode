// 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points

package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for _, val := range vals[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return []int{-1, -1}
	}

	// Extract the values from the linked list
	vals := []int{}
	for current := head; current != nil; current = current.Next {
		vals = append(vals, current.Val)
	}

	indecesOfCritPts := []int{}

	for i := 1; i < len(vals)-1; i++ {
		prev := vals[i-1]
		curr := vals[i]
		next := vals[i+1]

		if (curr > prev && curr > next) || (curr < prev && curr < next) {
			indecesOfCritPts = append(indecesOfCritPts, i)
		}
	}

	if len(indecesOfCritPts) < 2 {
		return []int{-1, -1}
	}

	minDistance := math.MaxInt64
	maxDistance := indecesOfCritPts[len(indecesOfCritPts)-1] - indecesOfCritPts[0]

	for i := 1; i < len(indecesOfCritPts); i++ {
		minDistance = min(minDistance, indecesOfCritPts[i]-indecesOfCritPts[i-1])
	}

	return []int{minDistance, maxDistance}
}

func main() {
	vals := []int{5, 3, 1, 2, 5, 1, 2}
	head := createLinkedList(vals)
	res := nodesBetweenCriticalPoints(head)
	fmt.Println("Result:", res)
}
