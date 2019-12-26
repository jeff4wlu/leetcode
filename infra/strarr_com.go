package infra

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
