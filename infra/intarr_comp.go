package infra

// 比较一个二维阵列是否一样
func IntArr2DComp(got, want [][]int) bool {

	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}

//比较两个[][]int集合内元素是否一样，而不是二维阵列。isSeq是判断子元素是否按顺序
func IntarrCollectionComp(a, b [][]int, isSeq bool) bool {

	if len(a) != len(b) {
		return false
	}

	dict := map[int][]int{}
	for i := 0; i < len(b); i++ {
		dict[i] = b[i]
	}
	for i := 0; i < len(a); i++ {
		idx := -1
		for k, v := range dict {
			if isSeq {
				if IntArrSeqCmp(a[i], v) {
					idx = k
				}
			} else {
				if IntArrNonSeqCmp(a[i], v) {
					idx = k
				}
			}
		}
		if idx == -1 {
			return false
		}
		delete(dict, idx)
	}

	if len(dict) != 0 {
		return false
	}

	return true
}

func compOneArr(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compOneString(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

//与顺序有关
func IntArrSeqCmp(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

//与顺序无关
func IntArrNonSeqCmp(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				found = true
			}
		}
		if !found {
			return false
		}

	}

	return true
}
