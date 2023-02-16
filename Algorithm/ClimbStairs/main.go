package main

import "math"

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢

func climbStairs(n int) int {
	sqrt:=math.Sqrt(5)
	powOne:=math.Pow((1+sqrt)/2,float64(n+1))
	powTwo:=math.Pow((1-sqrt)/2,float64(n+1))
	return int(math.Round((powOne-powTwo)/sqrt))
}