package leet187

//最直接的思路，可以优化为更加节省内存的方式
func RepeatDNA(s string) []string {

	n := len(s)
	res := []string{}
	if n < 11 {
		return res
	}
	dict := map[string]int{}
	for i := 0; i < n-10; i++ {
		dict[s[i:10+i]]++
		if dict[s[i:10+i]] == 2 {
			res = append(res, s[i:10+i])
		}
	}
	return res
}

//对上面算法的优化，使用位操作来减少字符串的存储大小。
//并且记录下已经转化的字母。
//以上算法成功的原因在于DNA只有4个字母表示，而且这4个字母的ASC码的最后三位都不一样。
func RepeatDNA1(s string) []string {

	n := len(s)
	res := []string{}
	if n < 11 {
		return res
	}
	dict := map[int]int{}
	var num int
	for i := 0; i < 10; i++ {
		num <<= 3
		num |= int(s[i]) & 7
	}
	dict[num]++
	num &= 0x7ffffff //保留最近9位的压缩版字母，不用每次都重新计算一遍。
	for i := 10; i < n; i++ {
		num <<= 3
		num |= int(s[i]) & 7
		dict[num]++
		if dict[num] == 2 {
			res = append(res, s[i-9:i+1])
		}
		num &= 0x7ffffff
	}
	return res
}
