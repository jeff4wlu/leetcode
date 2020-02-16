package leet267

//利用map，dfs找出不重复的排列组合。
func Palindrome(word string) []string {

	n := len(word)
	if n == 0 {
		return []string{}
	}

	dict := map[byte]int{}
	var mid byte
	halfArr := []byte{}
	for i := 0; i < n; i++ {
		dict[word[i]]++
	}
	for k, v := range dict {
		a := v / 2
		b := v % 2
		if b != 0 {
			if uint(mid) != 0 {
				return []string{} //不是回文
			}
			mid = k
		}
		for i := 0; i < a; i++ {
			halfArr = append(halfArr, k)
		}
	}

	n = len(halfArr)
	if n == 0 {
		return []string{string(mid)}
	}

	path := []byte{}
	res := [][]byte{}
	dfs(halfArr, path, &res)

	re := []string{}
	for j := 0; j < len(res); j++ {
		n = len(res[j])
		var tmp []byte
		if uint(mid) != 0 {
			tmp = make([]byte, 2*n+1)
		} else {
			tmp = make([]byte, 2*n)
		}
		for i := 0; i < n; i++ {
			if uint(mid) != 0 {
				tmp[i], tmp[2*n-i] = res[j][i], res[j][i]
			} else {
				tmp[i], tmp[2*n-i-1] = res[j][i], res[j][i]
			}
		}
		if uint(mid) != 0 {
			tmp[n] = mid
		}
		re = append(re, string(tmp))
	}

	return re
}

func dfs(arr []byte, path []byte, res *[][]byte) {

	n := len(arr)
	if n == 0 {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		(*res) = append((*res), tmp)
		return
	}

	dict := map[byte]int{}
	for i := 0; i < n; i++ {
		if _, ok := dict[arr[i]]; !ok {
			dict[arr[i]]++
			path = append(path, arr[i])
			tmp := []byte{}
			tmp = append(tmp, arr[:i]...)
			tmp = append(tmp, arr[i+1:]...)
			dfs(tmp, path, res)
			path = path[:len(path)-1]
		}
	}
}
