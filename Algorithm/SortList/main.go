package main

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链

type ListNode struct {
  Val int
  Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return sort(head,nil)
}

func merge(head1,head2 *ListNode)*ListNode{
	emptyHead:=&ListNode{}
	temp,temp1,temp2:=emptyHead,head1,head2
	for temp1!=nil&&temp!=nil{
		if temp1.Val<=temp2.Val{
			temp.Next=temp1
			temp1=temp1.Next
		}else{
			temp.Next=temp2
			temp2=temp2.Next
		}
		temp=temp.Next
	}

	if temp1!=nil{
		temp.Next=temp1
	}else if temp2!=nil{
		temp.Next=temp2
	}
	return emptyHead.Next
}

func sort(head,tail *ListNode)*ListNode{
	if head==nil{
		return head
	}
	if head.Next==tail{
		head.Next=nil
		return head
	}

	slow,fast:=head,head
	for fast!=nil{
		slow=slow.Next
		fast=fast.Next
		if fast!=tail{
			fast=fast.Next
		}
	}
	middle:=slow
	return merge(sort(head,middle),sort(middle,tail))
}