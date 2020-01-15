package leet115

func DistinctSeq(s, t string) int {
	path := []int{}
	res := [][]int{}
	helper(s, t, path, &res)
	return len(res)
}

func helper(s, t string, path []int, res *[][]int) {

	sLen := len(s)
	tLen := len(t)

	if tLen == 0 && sLen >= 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	if tLen == 0 || sLen == 0 && tLen != 0 || sLen < tLen {
		return
	}

	i := 0
	for i < sLen {
		for ; i < sLen; i++ {
			if t[0] == s[i] {
				break
			}
		}
		if i == sLen {
			return
		}
		path = append(path, i)
		helper(s[i+1:], t[1:], path, res)
		path = path[:len(path)-1]
		i++
	}

}
