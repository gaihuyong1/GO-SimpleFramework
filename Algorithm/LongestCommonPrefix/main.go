package main

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""

func longestCommonPrefix(strs []string) string {
	if len(strs)==0{
		return ""
	}
	res:=strs[0]
	for i:=1;i<len(strs);i++{
		j:=0
		for ;j<len(res)&&j<len(strs[i]);j++{
			if res[j]!=strs[i][j]{
				break
			}
			res=res[0:j]
			if res==""{
				return res
			}
		}
	}
	return res
}