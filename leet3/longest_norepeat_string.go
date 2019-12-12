package leet3

func LengthOfLongestSubstring(str []byte) int {

	hm := make(map[byte]int, 20)
	left := -1
	var length int

	for i, v := range str {

		if _, ok := hm[v]; ok {
			var l int
			if hm[v] > left { //滑动窗口内
				l = i - left - 1
				left = i
			} else {
				hm[v] = i
			}

			if l > length {
				length = l
			}
		}

		hm[v] = i

	}

	if length == 0 {
		length = len(str)
	}

	return length

}
