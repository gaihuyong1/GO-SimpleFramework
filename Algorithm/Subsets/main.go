package main

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集

func subsets(nums []int) [][]int {
	length:=len(nums)
	res:=make([][]int,0,1<<length)
	path:=make([]int,0,length)
	var dfs func(int)
	dfs=func(i int){
		res=append(res, append([]int(nil),path...))
		if i==length{
			return 
		}

		for j:=i;j<length;j++{
			path=append(path, nums[i])
			dfs(j+1)
			path=path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}