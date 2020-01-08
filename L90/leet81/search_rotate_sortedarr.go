package leet81

import "leetcode/infra"

//leet33的加强版
func SearchRotateSortedArr(in []int, target int) bool {

	return solve(in, target)

}

func solve(in []int, target int) bool {

	n := len(in)

	if n == 0 {
		return false
	}

	if n == 1 {
		if in[0] == target {
			return true
		}
		return false
	}

	mid := (n+1)/2 - 1 //mid是索引，0开始

	for cnt := 0; in[mid] == in[n-1]; cnt++ { //左移一位，直到不相等;并且要防止一直重复的数字
		tmp := in[0]
		for i := 1; i < n; i++ {
			in[i], in[i-1] = in[i-1], in[i]
		}
		in[n-1] = tmp
		if cnt > n-1 {
			if in[0] == target {
				return true
			}
			return false
		}
	}

	if in[mid] < in[n-1] { //右边有序
		if target == in[mid] {
			return true
		} else if target > in[mid] { //二分法
			if -1 != infra.BinarySearch(in[mid+1:], target) {
				return true
			}
			return false
		} else {
			return solve(in[:mid], target)
		}

	} else { //左边有序
		if target == in[mid] {
			return true
		} else if target < in[mid] { //二分法
			if -1 != infra.BinarySearch(in[:mid], target) {
				return true
			}
			return false
		} else {
			return solve(in[mid+1:], target)
		}
	}

}
