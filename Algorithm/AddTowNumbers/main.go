package main

// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

// 请你将两个数相加，并以相同形式返回一个表示和的链表。

// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode)(head *ListNode) {
	//创建空节点用来存储进位
	var endNode *ListNode
	carry:=0
	//依次计算两个链表的每一位结果的和
	for l1!=nil||l2!=nil{
		node1,node2:=0,0
		if l1!=nil{
			node1=l1.Val
			l1=l1.Next
		}
		if l2!=nil{
			node2=l2.Val
			l2=l2.Next
		}
		sum:=node1+node2+carry
		sum=sum%10
		carry=sum/10
		if head==nil{
			head=&ListNode{Val:sum}
			endNode=head
		}else{
			endNode.Next=&ListNode{Val:sum}
			endNode=endNode.Next
		}
	}
	if carry>0{
		endNode.Next=&ListNode{Val:carry}
	}
	return
}