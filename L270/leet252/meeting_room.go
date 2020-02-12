package leet252

func MeetingRoom(ranges [][]int) bool {

	times := len(ranges)
	if times == 1 {
		return true
	}

	var tmp int
	for i := 0; i < times; i++ {
		if ranges[i][0]-tmp < 0 {
			return false
		}
		tmp = ranges[i][1]
	}
	return true
}
