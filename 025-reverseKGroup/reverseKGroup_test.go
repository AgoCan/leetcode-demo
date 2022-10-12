package reversekgroupDemo

import (
	"fmt"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	a := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	b := reverseKGroup(a, 2)
	for b != nil {
		fmt.Println(b.Val)
		b = b.Next
	}
}
