package leet228

import "strconv"

import "fmt"

//利用题目的特点条件，第一是已排序好的数组，第二是连续的（即相差为一）
func SummaryRanges(nums []int) []string {

	res := []string{}
	n := len(nums)
	if n == 0 {
		return res
	}
	if n == 1 {
		res = append(res, getRanges(n, n))
		return res
	}

	var i, j int
	for j < n {
		if j-i != nums[j]-nums[i] {
			res = append(res, getRanges(nums[i], nums[j-1]))
			i = j
		} else {
			j++
		}
	}
	res = append(res, getRanges(nums[i], nums[j-1]))
	return res
}

func getRanges(s, e int) string {

	if s == e {
		return strconv.Itoa(s)
	}
	return fmt.Sprintf("%d->%d", s, e)

}
