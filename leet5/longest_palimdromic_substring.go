package leet5

func LongestPalimdromicSubstr(in string) string {
	var data []byte = []byte(in)
	n := len(data)

	// start与maxlen分别表示最长回文子串的起点跟长度
	var start, maxLen int
	for i := 0; i < n; {
		//遍历s中的字符的时候，我们首先判断剩余的字符数是否小于等于 maxLen 的一半，
		//是的话表明就算从当前到末尾到子串是半个回文串，那么整个回文串长度最多也就是 maxLen，
		//既然 maxLen 无法再变长了（无法比最长的长就应该停止了），计算这些就没有意义，
		//直接在当前位置 break 掉就行了
		if n-i <= maxLen/2 {
			break
		}
		left := i
		right := i
		//跳过重复字符，对奇偶的情况都可行
		for right < n-1 && data[right] == data[right+1] {
			right++
		}
		i = right + 1

		for left > 0 && right < n-1 && data[left-1] == data[right+1] {
			left--
			right++
		}

		if maxLen < right-left+1 {
			maxLen = right - left + 1
			start = left
		}

	}

	return string(data[start : start+maxLen])
}
