package leet28

func ImplementStrstr(haystack, needle string) int {

	if haystack == needle {
		return 0
	}

	hLen := len(haystack)
	nLen := len(needle)

	if hLen < nLen {
		return -1
	}

	var j int
	for i := 0; i < hLen; i++ {

		if j >= nLen {
			return i - nLen
		}
		if haystack[i:i+1] != needle[j:j+1] {
			j = 0
		} else {
			j++
		}

		if (hLen - i) < nLen {
			break
		}
	}
	return -1
}
