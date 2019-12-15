package leet5

//常规思维的算法
func LongestPalimdromicSubstr(in string) string {
	var data = []byte(in)
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

//使用动态规划思维dynamic programming
//每次判断的结果都依赖上一次判断的结果，这种问题的解答过程叫动态规划。
//把多阶段过程转化为一系列单阶段问题，利用各阶段之间的关系，逐个求解，
//创立了解决这类过程优化问题的新方法——动态规划.简单的说就是通过分阶段求解问题的结果，
//每一阶段的结果都依赖于上一阶段的结果。
func LongestPalimdromicSubstr1(in string) string {
	var data = []byte(in)
	n := len(data)

	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	var start, len int

	//[i,j]代表回文字符串，在二维数组中就是下半三角区
	for i := 0; i < n; i++ {

		dp[i][i] = 1

		for j := 0; j < n; j++ {

			if data[i] == data[j] && (i-j < 2 || dp[j+1][i-1] == 1) {
				dp[j][i] = 1
				if len < i-j+1 {
					len = i - j + 1
					start = j
				}
			}

		}

	}

	return string(data[start : start+len])
}
