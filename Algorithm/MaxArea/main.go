package main

// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 返回容器可以储存的最大水量

//使用双指针，从左右两边向中间移动
func maxArea(height []int) int {
  left,right,res:=0,len(height)-1,0
	for left<right{
		if height[left]<height[right]{
			res=max(res,height[left]*(right-1))
			left+=1
		}else{
			res=max(res,height[right]*(right-1))
			right-=1
		}
	}
	return res
}

func max(x,y int)int{
	if x<y{
		return y
	}
	return x
}