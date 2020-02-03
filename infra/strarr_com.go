package infra

//与顺序无关
func StringArrCollectionComp(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if StringArrComp(av, bv) {
				found = true
				break
			}
		}

		if !found {
			return false
		}

	}
	return true
}

//与顺序无关
func StringArrComp(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}
	dict := map[string]int{}
	for _, v := range a {
		dict[v]++
	}
	for _, v := range b {
		if _, ok := dict[v]; !ok {
			return false
		}
		dict[v]--
	}

	for _, v := range a {
		if dict[v] != 0 {
			return false
		}
	}

	return true
}

//与顺序有关
func StringArrSeqComp(a, b []string) bool {

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
