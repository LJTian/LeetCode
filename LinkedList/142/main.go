package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将节点设置为key 如果存在 可以获取到值
func detectCycle(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return head
		}
		seen[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 还有一种方式需要计算公式计算
