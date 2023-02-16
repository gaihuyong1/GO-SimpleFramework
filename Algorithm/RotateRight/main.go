package main

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置

func rotateRight(head *ListNode, k int) *ListNode {
 if head==nil||head.Next==nil||k==0{
	return head
 }

 //获取链表长度
 len:=1
 temp:=head
 for temp.Next!=nil{
	temp=temp.Next
	len++
 }

 //计算需要移动的数量
 move:=len-k%len
 if move==len{
	return head
 }
 temp.Next=head
 if move>0{
	temp=temp.Next
	move--
 }
 res:=temp.Next
 temp.Next=nil
 return res
}

type ListNode struct {
  Val int
  Next *ListNode
}