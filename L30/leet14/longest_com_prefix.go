package leet14

func minString(ss []string) int {
	res := len(ss[0])
	for i := 1; i < len(ss); i++ {
		if len(ss[i]) < res {
			res = len(ss[i])
		}
	}
	return res

}

// 字符串切片，a[0:]即使是空串也不会出错，但a[0:1]的空串就会出错
func LongestComPrefix(ss []string) string {

	if len(ss) < 2 {
		return ""
	}

	//求出最短字符串长度
	minLen := minString(ss)
	if minLen == 0 {
		return ""
	}

	end := minLen

	for i := 0; i < minLen; i++ {

		//第一个字符串作为比较的基准
		compare := ss[0][i : i+1]

		if compare == "" {
			break
		}

		equal := true
		for j := 1; j < len(ss); j++ {
			if compare != ss[j][i:i+1] {
				equal = false
				break
			}
		}

		if !equal {
			end = i
			break
		}

	}

	return ss[0][:end]

}
