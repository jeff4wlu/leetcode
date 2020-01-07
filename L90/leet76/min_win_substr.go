package leet76

//指针对碰，滑动窗口
func MinWinSubstr(s, t string) string {

	sLen := len(s)
	tLen := len(t)

	var res string
	var left, cnt int
	min := sLen

	dict := map[string]int{} //分配空内存
	for i, _ := range t {
		dict[t[i:i+1]]++
	}

	for i := 0; i < sLen; i++ {
		if _, ok := dict[s[i:i+1]]; ok {
			dict[s[i:i+1]]--
			if dict[s[i:i+1]] >= 0 {
				cnt++
			}
		}
		for cnt == tLen {
			if min > i-left+1 {
				min = i - left + 1
				res = s[left : i+1]
			}
			if _, ok := dict[s[left:left+1]]; ok {
				dict[s[left:left+1]]++
				if dict[s[left:left+1]] > 0 {
					cnt--
				}
			}
			left++
		}

	}

	return res
}
