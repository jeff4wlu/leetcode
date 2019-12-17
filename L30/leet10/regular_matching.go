package leet10

//可以对string做切片操作，出来的也是string
//string的内存表示为两个字，一个字为指针，第二为长度
//string[len:]为空字符串
func RegularMatching(s, p string) bool {

	//size等于0的情况
	if s == "" {
		if p == "" {
			return s == ""
		}
	}

	//size等于1的情况
	if len(p) == 1 {
		return (len(s) == 1 && (s[0:1] == p[0:1] || p[0:1] == "."))
	}

	//size大于1的情况，且不能*
	if p[1:2] != "*" {
		if s == "" {
			return false
		}
		return (s[0:1] == p[0:1] || p[0:1] == ".") && RegularMatching(s[1:], p[1:])
	}

	//处理a*或.*模式时的情况
	for s != "" && (s[0:1] == p[0:1] || p[0:1] == ".") {
		s = s[1:]
	}

	return RegularMatching(s, p[2:])

}
