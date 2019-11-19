package main

import (
	"github.com/mats9693/leetcode/utils/compare"
	"testing"
)

type In struct {
	Node *compare.ListNode
	N    int
}

var testCase = []struct {
	In     In
	Except *compare.ListNode
}{
	{In{compare.MakeList(1, 2, 3, 4, 5), 2}, compare.MakeList(1, 2, 3, 5)},
	{In{compare.MakeList(1), 1}, nil},
	{In{compare.MakeList(1, 2), 2}, compare.MakeList(2)},
}

func TestRemoveNthFromEnd(t *testing.T) {
	tcs := testCase
	for i := range tcs {
		if !compare.CompareOnList(removeNthFromEnd(tcs[i].In.Node, tcs[i].In.N), tcs[i].Except) {
			t.Errorf("remove nth from end test failed on case: %d", i)
		}
	}
}
