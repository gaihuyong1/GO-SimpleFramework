package main

import "sort"

// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

// 你返回所有和为 0 且不重复的三元组

func threeSum(nums []int) [][]int {
	len:=len(nums)
	sort.Ints(nums)
	res:=make([][]int,0)
	if len<3{
		return res
	}
	for i:=0;i<len;i++{
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		j:=len-1
		target:=-1*nums[i]
		for k:=i+1;k<len;k++{
			if j<i+1&&nums[k]==nums[k-1]{
				continue
			}
			for k<j&&nums[k]+nums[j]>target{
				j--
			}

			if k==j{
				break
			}
			if nums[k]+nums[j]==target{
				res=append(res,[]int{nums[i],nums[k],nums[j]})
			}
		}
	}
	return res
}