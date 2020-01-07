package leet78

func Subsets(sets []int) [][]int {

	res := [][]int{}
	path := []int{}

	for i := 0; i < len(sets); i++ {
		for j := 0; j < len(sets); j++ {
			solve(sets[j], i+1, sets[j+1:], path, &res)
		}

	}
	return append(res, []int{})
}

func solve(el, length int, rest, path []int, res *[][]int) {

	pLen := len(path) + 1

	if pLen > length {
		return
	}

	if pLen == length {
		path = append(path, el)
		tmp := make([]int, pLen)
		copy(tmp, path)
		(*res) = append((*res), tmp)
		return
	}

	path = append(path, el)
	for i := 0; i < len(rest); i++ {
		solve(rest[i], length, rest[i+1:], path, res)
	}

}
