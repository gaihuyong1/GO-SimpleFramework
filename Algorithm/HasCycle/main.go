package main

// 给你一个链表的头节点 head ，判断链表中是否有环。

// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。

// 如果链表中存在环 ，则返回 true 。 否则，返回 false


 type ListNode struct {
    Val int
    Next *ListNode
}

func hasCycle(head *ListNode) bool {
	record:=map[*ListNode]struct{}{}
	for head!=nil{
		if _,ok:=record[head];ok{
			return true
		}
		record[head]=struct{}{}
		head=head.Next
	}

	return false
}