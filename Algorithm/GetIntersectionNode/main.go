package main

// 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null

type ListNode struct {
  Val int
  Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA==nil||headB==nil{
			return nil
		}
		nodeA,nodeB:=headA,headB
		for nodeA!=nodeB{
			if nodeA==nil{
				nodeA=nodeB
			}else{
				nodeA=nodeA.Next
			}
			if nodeB==nil{
				nodeB=nodeA
			}else{
				nodeB=nodeB.Next
			}
		}
		return nodeA
}