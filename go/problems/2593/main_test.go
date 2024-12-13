package main

import (
	"context"
	"testing"
	"time"
)

type testCase struct {
	nums     []int
	expected int64
}

type testResult struct {
	got      int64
	expected int64
}

func TestFindScore(t *testing.T) {
	testCases := []testCase{
		{[]int{2, 1, 3, 4, 5, 2}, 7},
		{[]int{2, 3, 5, 1, 3, 2}, 5},
		{[]int{4, 3, 2, 2}, 6},
		{[]int{10, 44, 10, 8, 48, 30, 17, 38, 41, 27, 16, 33, 45, 45, 34, 30, 22, 3, 42, 42}, 212},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			resultChan := make(chan testResult)

			go func() {
				got := findScore(tc.nums)
				resultChan <- testResult{got, tc.expected}
			}()

			select {
			case result := <-resultChan:
				if result.got != result.expected {
					t.Errorf("got %v, expected %v", result.got, result.expected)
				}
			case <-ctx.Done():
				t.Error("time limit exceeded")
			}
		})
	}
}
