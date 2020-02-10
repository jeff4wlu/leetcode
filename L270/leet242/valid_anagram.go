package leet242

//根据题目特点26个字母，使用array来模拟hashmap
func ValidAnagram(s, t string) bool {

	ls := len(s)
	lt := len(t)
	if ls != lt {
		return false
	}

	hmap := make([]int, 26)
	for i := 0; i < ls; i++ {
		hmap[int(s[i]-'a')]++
		hmap[int(t[i]-'a')]--
	}

	for i := 0; i < ls; i++ {
		if hmap[i] != 0 {
			return false
		}
	}

	return true
}
