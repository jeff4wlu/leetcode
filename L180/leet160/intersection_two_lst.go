package leet160

import "leetcode/infra"

func IntersectionLst(a, b *infra.ListNode) *infra.ListNode {

	var aLen, bLen int
	ap := a
	bp := b

	for ap != nil {
		aLen++
		ap = ap.Next
	}
	for bp != nil {
		bLen++
		bp = bp.Next
	}

	ap = a
	bp = b

	if aLen > bLen {
		tmp := aLen - bLen
		for tmp != 0 {
			tmp--
			ap = ap.Next
		}
	} else {
		tmp := bLen - aLen
		for tmp != 0 {
			tmp--
			bp = bp.Next
		}
	}

	for ap != bp {
		ap = ap.Next
		bp = bp.Next
	}

	return ap
}
