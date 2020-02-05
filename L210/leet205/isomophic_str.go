package leet205

//需要使用两个hashmap，因为只用一个map会导致多个字符映射为一个，会出错，参加测试用例1
func IsomophicStr(a, b string) bool {

	al := len(a)
	bl := len(b)
	if al != bl {
		return false
	}

	d1, d2 := map[string]string{}, map[string]string{}

	var s1, s2 string
	for i := 0; i < al; i++ {
		if _, ok := d1[a[i:i+1]]; !ok {
			d1[a[i:i+1]] = b[i : i+1]
		}
		s1 += d1[a[i:i+1]]
		if _, ok := d2[b[i:i+1]]; !ok {
			d2[b[i:i+1]] = a[i : i+1]
		}
		s2 += d2[b[i:i+1]]

	}
	if s1 != b || s2 != a {
		return false
	}

	return true

}
