package reversekgroupDemo

import "fmt"

/*

> 简单的描述：  5个元素，k等于2 就反转 12 34 其中5不变。  5个元素，k等于3 就反转 123 其中45不变

给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

1 -> 2 -> 3 -> 4 -> 5
          |
2 -> 1 -> 4 -> 3 -> 5

示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]


1 -> 2 -> 3 -> 4 -> 5
          |
3 -> 2 -> 1 -> 4 -> 5

示例 2：
输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	fmt.Println(head, tail)
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	fmt.Println(head, tail)
	return tail, head
}
