package main


// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0


func reverse(x int) (num int) {
	last:=0
	for x!=0{
		temp:=x%10
		last=num
		num=num*10+temp
		if(last!=num/10){
			return 0
		}
		x/=10
	}
	return num
}