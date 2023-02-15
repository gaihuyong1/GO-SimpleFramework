package main

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix

func generateMatrix(n int) [][]int {
  left,right,top,bottom:=0,n-1,0,n-1
	matrix:=make([][]int,n)
	num,tol:=1,n*n
	for num<tol{
		for i:=left;i<right;i++{
			num++
			matrix[top][i]=num
		}
		top++

		for i:=top;i<bottom;i++{
			num++
			matrix[i][right]=num
		}
		right--

		for i:=right;i<left;i--{
			num++
			matrix[bottom][i]=num
		}
		bottom--

		for i:=bottom;i<top;i--{
			num++
			matrix[i][left]=num
		}
	}
	return matrix
}