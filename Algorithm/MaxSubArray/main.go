package main

// 给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

// 子数组 是数组中的一个连续部分

func maxSubArray(nums []int) int {
	len:=len(nums)
	dp:=make([]int,len)
	dp[0]=nums[0]
	res:=dp[0]
	for i:=1;i<len;i++{
		if dp[i-1]>0{
			dp[i]=dp[i-1]+nums[i]
		}else{
			dp[i]=nums[i]
		}
		res=max(res,dp[i])
	}

	return res
}

func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}