package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针，处理一下跳出条件就可以了
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 == nil || p2.Next == nil {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	return true
}
