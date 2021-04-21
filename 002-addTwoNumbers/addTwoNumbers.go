package addTwoNumbersDemo

// 还不会，过后再看

/*

中等

给出两个 非空 的链表用来表示两个非负的整数。
其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

type listNode struct {
	Val  int
	Next *listNode
}

func AddTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	s := new(listNode) // &{0 <nil>} 所以必须有值
	p := s             //
	u := 0             // 进1位置， 例如9+3=12， 进一位为 1
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		p.Val = (l1.Val + l2.Val + u/10) % 10 // 先算个位数，9 + 8 = 17 求余
		u = l1.Val + l2.Val + u/10 - p.Val    // u = 1
		if l1.Next != nil && l2.Next != nil {
			p.Next = new(listNode) // 如果有下一位数字，则继续开辟一个空间
			p = p.Next
		}
	}
	// 判断，如果参数都是为nil，就已经可以直接退出了。在第一行
	if l1 == nil && l2 == nil && u == 0 {
		// 没有任何
		return s
	} else if l1 == nil && l2 == nil {
		// 同为3位数，然后都加完了。
		p.Next = new(listNode)
	} else if l1 == nil {
		// 下面两个把剩余的数字给加上
		p.Next = l2
	} else if l2 == nil {
		p.Next = l1
	}
	p = p.Next
	for tmp := 0; u != 0 && p != nil; p = p.Next {
		tmp = p.Val + u/10
		p.Val = (p.Val + u/10) % 10
		u = tmp - p.Val
		if p.Next == nil && u != 0 {
			p.Next = new(listNode)
			p.Next.Val = u / 10
			break
		}
	}
	return s
}
