package main

import (
	"math"
	"sort"
)

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。

// 返回这三个数的和。

// 假定每组输入只存在恰好一个解

func threeSumClosest(nums []int, target int) int {
  sort.Ints(nums)
	res:=nums[0]+nums[1]+nums[2]
	for i:=0;i<len(nums);i++{
		start:=i+1
		end:=len(nums)-1
		for start<end{
			sum:=nums[i]+nums[start]+nums[end]
			if math.Abs(float64(target-sum))<math.Abs(float64(target-res)){
				res=sum
			}
			if sum>target{
				end--
			}else if sum<target{
				start++
			}else{
				return res
			}
		}
	}
	return res
}
