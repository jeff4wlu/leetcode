package infra

// 比较一个二维阵列是否一样
func IntArr2DComp(got, want [][]int) bool {

	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[0]); j++ {
			if got[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}

//比较两个[][]int集合内元素是否一样，而不是二维阵列
func IntarrCollectionComp(a, b [][]int) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		isSame := false
		for j := 0; j < len(b); j++ {
			if compOneArr(a[i], b[j]) {
				isSame = true
				break
			}
		}
		if !isSame {
			return false
		}
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
