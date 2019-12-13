package leet4

import "leetcode/base"

//0、保证arr1比arr2短
//1、i是短数组的Index，j是长数组的index。
//2、min是短数组的开始，max是短数组的结束
//3、无论n+m是偶数还是基数，都可归结为找(m+n)/2，之后再按基偶来做计算
//4、找中位数就是找出划分两个数组合适的边界，但要注意左右都有可能分出空的分段。
func MedianTwoSortedArr(arr1, arr2 []int) float64 {

	var minIndex, maxIndex, i, j, len1, len2, median int

	if len(arr1) > len(arr2) {
		arr1, arr2 = arr2, arr1
	}

	len1 = len(arr1)
	len2 = len(arr2)

	maxIndex = len1

	for {

		if minIndex > maxIndex {
			break
		}

		//i是短的index
		i = (minIndex + maxIndex) / 2
		j = (len1+len2+1)/2 - i

		if i < len1 && j > 0 && arr2[j-1] > arr1[i] {
			minIndex = i + 1
		} else if i > 0 && j < len2 && arr1[i-1] > arr2[j] {
			maxIndex = i - 1
		} else {
			if i == 0 {
				median = arr2[j-1]
			} else if j == 0 {
				median = arr2[i-1]
			} else {
				median = base.IntMax(arr1[i-1], arr2[j-1])
			}

			break

		}

	}

	// calculating the median.
	// If number of elements is odd
	// there is one middle element.
	if (len1+len2)%2 == 1 {
		return float64(median)
	}

	// Elements from a[] in the
	// second half is an empty set.
	if i == len1 {
		return float64(median+arr2[j]) / 2
	}

	// Elements from b[] in the
	// second half is an empty set.
	if j == len2 {
		return float64(median+arr1[i]) / 2
	}

	return float64(median + base.IntMin(arr1[i], arr2[j])/2)

}
