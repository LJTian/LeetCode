package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1：对比合并 迭代
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {

	pevr := new(ListNode)

	p1 := list1
	p2 := list2
	node := pevr

	for p1 != nil || p2 != nil {

		if p1 == nil {
			pevr.Next = p2
			break
		}
		if p2 == nil {
			pevr.Next = p1
			break
		}
		if p1.Val < p2.Val {
			pevr.Next = p1
			p1 = p1.Next
		} else {
			pevr.Next = p2
			p2 = p2.Next
		}
		pevr = pevr.Next
	}

	return node.Next
}

// 递归
// 递归，找到结束标志，将每次递归的返回值都当作是已经完成的结果，思考完成的结果与当前的处理逻辑，写出来就可以了
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	// 如果有一个链表为nil 则剩下的就为另一个链表
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	// 找到小的值，将小的第二个节点再此比较
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list2.Next, list1)
		return list2
	}
}
