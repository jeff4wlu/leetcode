package leet247

var dict []byte = []byte{'1', '6', '8', '9', '6', '8', '9', '1'}

func StrobogrammaticNum(k int) []string {

	path := make([]byte, k)
	res := []string{}
	dfs(k, 0, path, &res)
	return res
}

//step是0base
func dfs(k, step int, path []byte, res *[]string) {

	mid := (k - 1) / 2 //index

	if step == mid+1 {
		(*res) = append((*res), string(path))
		return
	}

	if k%2 != 0 && step == mid {
		tmp := make([]byte, k)
		copy(tmp, path)
		tmp[step] = '1'
		dfs(k, step+1, tmp, res)
		tmp = make([]byte, k)
		copy(tmp, path)
		tmp[step] = '8'
		dfs(k, step+1, tmp, res)

	} else {
		for i := 0; i <= 3; i++ {
			tmp := make([]byte, k)
			copy(tmp, path)
			tmp[step] = dict[i]
			tmp[k-1-step] = dict[7-i]

			dfs(k, step+1, tmp, res)
		}
	}

}
