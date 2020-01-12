package infra

//与顺序无关
func StringArrCollectionComp(a, b [][]string) bool {
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

func StringArrComp(a, b []string) bool {
	if len(a) == 0 || len(b) == 0 {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
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
