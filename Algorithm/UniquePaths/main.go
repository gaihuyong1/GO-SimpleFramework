package main

// 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。

// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。

// 问总共有多少条不同的路径

func uniquePaths(m int, n int) int {
	dp:=make([][]int,m)
	for i:=range dp{
		dp[i]=make([]int,n)
		dp[i][0]=1
	}

	for in:=0;in<n;in++{
		dp[0][in]=1
	}

	for im:=1;im<m;im++{
		for j:=1;j<n;j++{
			dp[im][j]=dp[im-1][j]+dp[im][j-1]
		}
	}
	return dp[m-1][n-1]
}