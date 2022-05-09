package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail, tmp *ListNode
	num, carry := 0, 0

	// 跳出条件
	for l1 != nil || l2 != nil {

		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		// 计算值
		sum := v1 + v2 + carry
		num, carry = sum%10, sum/10

		// 看一下是否是第一个节点
		if tail == nil {
			tmp = &ListNode{num, nil}
			tail = tmp
		} else {
			tmp.Next = &ListNode{num, nil}
			tmp = tmp.Next
		}
	}

	// 如果有进位需要再挂一个节点
	if carry > 0 {
		tmp.Next = &ListNode{1, nil}
	}

	return tail
}
