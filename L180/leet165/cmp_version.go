package leet165

import "strconv"

func CmpVersion(a, b string) int {

	alen := len(a)
	blen := len(b)

	var apos, bpos int
	for {
		if apos >= alen && bpos < blen {
			return -1
		}
		if bpos >= blen && apos < alen {
			return 1
		}

		anode, bnode := []byte{}, []byte{}
		for ; apos < alen && a[apos] != '.'; apos++ {
			anode = append(anode, a[apos])
		}
		for ; bpos < blen && b[bpos] != '.'; bpos++ {
			bnode = append(bnode, b[bpos])
		}

		anum, _ := strconv.Atoi(string(anode))
		bnum, _ := strconv.Atoi(string(bnode))
		if anum > bnum {
			return 1
		} else if anum < bnum {
			return -1
		}

		apos++
		bpos++
	}

	return 0
}
