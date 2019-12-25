package infra

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
