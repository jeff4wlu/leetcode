package leet77

//穷举，使用递归。关键点是这个组合是有序的，小在前。
func Combinations(n, k int) [][]int {

	rest := []int{}
	for i := 1; i <= n; i++ {
		rest = append(rest, i)
	}
	res := [][]int{}
	path := []int{}

	for i := 1; i < n; i++ {
		solve(i, k, rest[i:], path, &res)
	}
	return res
}

func solve(el, k int, rest, path []int, res *[][]int) {

	lvl := len(path) + 1

	if lvl > k {
		return
	}

	if lvl == k {
		path = append(path, el)
		tmp := make([]int, len(path))
		copy(tmp, path)
		(*res) = append((*res), tmp)
		return
	}

	path = append(path, el)
	for i, v := range rest {
		solve(v, k, rest[i+1:], path, res)
	}

}
