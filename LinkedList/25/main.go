package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1、记录头结点前节点，和尾结点的后节点
// 2、翻转选中的链表
// 3、链接前后链表

// 4、没有写出来

// 翻转链表
func reversalList(start *ListNode) *ListNode {

	if start == nil || start.Next == nil {
		return start
	}
	newNode := reversalList(start.Next)
	start.Next.Next = start
	start.Next = nil
	return newNode
}

//
//func reverseKGroup(head *ListNode, k int) *ListNode {
//
//	prev := &ListNode{0, head} // 头结点的前节点
//	backNode := head           // 尾结点的后一个节点
//	tailNode := head           // 头结点
//	endNode := prev            // 尾结点
//
//	for endNode != nil && endNode.Next != nil {
//
//		endNode = endNode.Next
//
//		for i := 0; i < k; i++ {
//
//		}
//		if endNode.Next != nil {
//
//		}
//
//		backNode = endNode.Next // 保存下一段节点头
//		//newNode := reversalList(tailNode, endNode) // 翻转链表
//		prev.Next = newNode
//		tailNode.Next = backNode
//
//		prev = tailNode
//		tailNode = backNode
//	}
//
//	return head
//}

func main() {

	array := []int{1, 2, 3, 4, 5}

	head := new(ListNode)

	p := head
	for _, v := range array {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	head = head.Next

	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}

	newNode2 := reversalList(head)
	for newNode2 != nil {
		fmt.Println(newNode2.Val)
		newNode2 = newNode2.Next
	}
	//
	//newHead := reverseKGroup(head, 2)
	//
	//for newHead != nil {
	//	fmt.Println(newHead.Val)
	//	newHead = newHead.Next
	//}

}
