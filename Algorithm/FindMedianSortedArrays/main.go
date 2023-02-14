package main

import "math"

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 算法的时间复杂度应该为 O(log (m+n))

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
  len1:=len(nums1)
  len2:=len(nums2)
	if len1>len2{
		return findMedianSortedArrays(nums2,nums1)
	}
	var lMax1,lMax2,rMin1,rMin2,c1,c2 int
	lo:=0
	hi:=2*len1
	for lo<=hi{
		c1 = (lo + hi) / 2;  //c1是二分的结果
			c2 = len2 + len1 - c1;

			if c1==0{
				lMax1=math.MinInt
			}else{
				lMax1= nums1[(c1 - 1) / 2]
			}

			if c1==2*len1{
				rMin1=math.MaxInt
			}else{
				rMin1=nums1[c1 / 2]
			}

			if c2==0{
				lMax2=math.MinInt
			}else{
				lMax2=nums2[(c2 - 1) / 2]
			}

			if c2==2*len2{
				rMin2=math.MinInt
			}else{
				rMin2=nums2[c2 / 2]
			}

			if lMax1 > rMin2{
				hi = c1 - 1;
			}else if lMax2 > rMin1{
				lo = c1 + 1;
			}else{
				break
			}
	}
	return (float64)(max(lMax1, lMax2) + min(rMin1, rMin2)) / 2
}


func min(x,y int) int{
 if x<y{
	return x
 }
 return y
}

func max(x,y int)int{
	if x<y{
	return y
 }
 return x
}