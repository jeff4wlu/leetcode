package leet27

func RemoveElement(in []int, num int) int {

	size := len(in)

	if size < 2 {
		if size == 0 {
			return 0
		} else if in[0] == num {
			return 0
		} else {
			return 1
		}
	}

	var a int
	for b := a; b < size; b++ {
		if in[b] != num {
			in[a], in[b] = in[b], in[a]
			a++
		}
	}

	return a
}
