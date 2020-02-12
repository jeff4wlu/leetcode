package leet252

//注意，题意是给出的时间段没有排序的。所以结束后要再和第一个比一下。
func MeetingRoom(ranges [][]int) bool {

	times := len(ranges)
	if times == 1 {
		return true
	}

	start := ranges[0][0]
	end := ranges[0][1]
	for i := 1; i < times; i++ {
		if !IsOverlap([]int{start, end}, ranges[i]) {
			return false
		}
		start, end = ranges[i][0], ranges[i][1]
	}

	if !IsOverlap([]int{start, end}, ranges[0]) {
		return false
	}

	return true
}

func IsOverlap(a, b []int) bool {

	if a[0] == b[0] {
		return false
	} else if a[0]-b[0] < 0 && a[1]-b[0] > 0 {
		return false
	} else if a[0]-b[0] > 0 && a[0]-b[1] < 0 {
		return false
	}
	return true
}
