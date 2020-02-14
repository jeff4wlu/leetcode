package leet253

import "leetcode/infra"

//线段重叠数的最大值就是需要的最小房间数。五个线段中有三个重叠在一起，就需要三个房间。
//把开始时间标记为1，结束时间标记为-1，然后把这些数据从小到大排列一起，计算出加起来的最大值。
//使用两个排序了的一维数组计算。
func MeetingRoom(intervals [][]int) int {

	var res int
	starts, ends := []int{}, []int{}
	for _, v := range intervals {
		starts = append(starts, v[0])
		ends = append(ends, v[1])
	}
	//从小到大
	starts = infra.BubbleSort(starts, false)
	ends = infra.BubbleSort(ends, false)

	var j int
	n := len(starts)
	for i := 0; i < n && j < n; {
		if ends[j] > starts[i] {
			res++
			i++
		} else {
			res--
			j++
		}
	}

	return res
}
