package main

import "math"

// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

// 路径和 是路径中各节点值的总和。

// 给你一个二叉树的根节点 root ，返回其 最大路径和

 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum:=math.MaxInt32

	var dfs func(root *TreeNode)int
	dfs=func(root *TreeNode)int{
		if root==nil{
			return 0
		}
		left:=dfs(root.Left)
		right:=dfs(root.Right)

		tempSum:=left+root.Val+right
		maxSum=max(maxSum,tempSum)
		res:=root.Val+max(left,right)
		return max(res,0)
	}
	dfs(root)
	return maxSum
}

func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}