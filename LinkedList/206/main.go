package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func reverseList(head *ListNode) *ListNode {

	var prev *ListNode // 遍历指针
	node := head       // 操作指针
	for node != nil {
		next := node.Next // 记录下一个节点，防止丢失
		node.Next = prev  // 翻转
		prev = node       // 指针偏移
		node = next       // 指针偏移
	}
	return prev
}

// 递归
func reverseList1(head *ListNode) *ListNode {

	// 寻找最后节点
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next) // 1->2   2->1
	head.Next.Next = head             // 1->2->1
	head.Next = nil                   // 1->null

	return newHead
}
