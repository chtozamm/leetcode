package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testCase struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

var testCases = []testCase{
	{
		l1:       listFromArray([]int{2, 4, 3}),
		l2:       listFromArray([]int{5, 6, 4}),
		expected: listFromArray([]int{7, 0, 8}),
	},
}

func TestSolution(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := addTwoNumbers(tc.l1, tc.l2)
			if !reflect.DeepEqual(arrayFromList(got), arrayFromList(tc.expected)) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
