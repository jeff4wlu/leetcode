package leet161

func OneEditDis(s, t string) bool {

	slen := len(s)
	tlen := len(t)
	if slen == tlen {
		var count int
		for i := 0; i < slen; i++ {
			if s[i] != t[i] {
				count++
			}
		}
		if count > 1 {
			return false
		}
		return true
	}

	dict := map[byte]int{}
	if tlen > slen {
		t, s = s, t
		tlen, slen = slen, tlen
	}
	if slen-tlen == 1 {
		var i int
		dict[s[i]]++
		i = 1
		for ; i < slen; i++ {
			dict[s[i]]++
			dict[t[i-1]]++
			if v, ok := dict[t[i-1]]; ok {
				if v != 2 {
					return false
				}
			}
		}
		return true

	}

	return false
}
