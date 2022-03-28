package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head

	return newHead
}

// 迭代
func swapPairs1(head *ListNode) *ListNode {

	dumpNode := &ListNode{0, head}
	tmpNode := dumpNode

	for tmpNode.Next != nil && tmpNode.Next.Next != nil {
		node1 := tmpNode.Next
		node2 := tmpNode.Next.Next

		tmpNode.Next = node2    // 0->2
		node1.Next = node2.Next // 1->3
		node2.Next = node1      // 2->1
		tmpNode = node1         // 后移处理下一个节点
	}

	return dumpNode.Next
}

// 创建哑结点-> 1->2->3   变换为 哑结点-> 2 ->1 ->3
