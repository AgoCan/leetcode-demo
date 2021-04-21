package addTwoNumbersDemo

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var l1 = new(listNode)
	var l2 = new(listNode)
	l1 = &listNode{
		Val: 2,
		Next: &listNode{
			Val: 4,
			Next: &listNode{
				Val: 3,
				Next: &listNode{
					Val:  0,
					Next: nil,
				},
			},
		},
	}
	l2 = &listNode{
		Val: 5,
		Next: &listNode{
			Val: 6,
			Next: &listNode{
				Val: 4,
				Next: &listNode{
					Val:  0,
					Next: nil,
				},
			},
		},
	}
	s := AddTwoNumbers(l1, l2)
	fmt.Println(s.Next.Next.Next.Val, s.Next.Next.Val, s.Next.Val, s.Val)
}
