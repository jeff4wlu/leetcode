package leet80

func RemoveDuplicate(in []int) int {

	n := len(in)
	var size int
	cnt := 1
	tmp := -1

	for i := 0; i < n; i++ {
		if tmp != in[i] {
			cnt = 1
			size++
			tmp = in[i]
		} else {
			cnt--
			if cnt == 0 {
				size++
			}
		}
	}

	return size
}
