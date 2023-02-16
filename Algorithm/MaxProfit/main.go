package main

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0

func maxProfit(prices []int) int {
	length:=len(prices)
	if length==0{
		return 0
	}
	dp:=make([]int,length)
	minPrice:=prices[0]

	for i:=1;i<length;i++{
		minPrice=min(minPrice,prices[1])
		dp[i]=max(dp[i-1],prices[i]-minPrice)
	}
	return dp[length-1]
}

func max(x,y int)int{
	if x>y{
		return x
	}
	return y
}

func min(x,y int)int{
	if x<y{
		return x
	}
	return y
}