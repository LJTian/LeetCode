package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 第一种答案：
	// 将A链表放置map中
	// 遍历B链表，发现存在后续则都相同，此节点即为相交节点

	// map1 := map[*ListNode]bool{}

	// for tmpA := headA; tmpA != nil; tmpA=tmpA.Next {
	//     map1[tmpA]=true
	// }

	// for tmpB := headB; tmpB != nil; tmpB=tmpB.Next{
	//     if _,ok := map1[tmpB]; ok {
	//         return tmpB
	//     }
	// }

	// return nil

	// 第二种答案：
	// 公式：  a + c + b = b + c + a
	// 其中：  a: 表示 A链表在相交前的长度
	//        b: 表示 B链表在相交前的长度
	//        c: 表示 相交后的长度

	tmpA, tmpB := headA, headB

	// 只要有一种为nil，肯定不会相交
	if tmpA == nil || tmpB == nil {
		return nil
	}

	for tmpA != tmpB {

		if tmpA == nil {
			tmpA = headB // 移到最后时，将指针指向B链
		} else {
			tmpA = tmpA.Next
		}
		if tmpB == nil {
			tmpB = headA // 移到最后时，将指针指向A链
		} else {
			tmpB = tmpB.Next
		}
	}
	return tmpA
}
