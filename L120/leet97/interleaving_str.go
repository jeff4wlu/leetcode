package leet97

//dfs,穷举回溯所有可能，即要考虑所有路径，本例相当于四叉树
func InterLeavingStr(s1, s2, s3 string) bool {

	return dfs(s1, s2, s3)

}

func dfs(s1, s2, s3 string) bool {
	s1L := len(s1)
	s2L := len(s2)
	s3L := len(s3)

	if s1L+s2L != s3L {
		return false
	}

	if s1L == 0 && s2 == s3 || s2L == 0 && s1 == s3 {
		return true
	}

	if s1L == 0 || s2L == 0 {
		return false
	}

	if s1[0] == s3[0] && s2[0] != s3[0] {
		return dfs(s1[1:], s2, s3[1:])
	} else if s2[0] == s3[0] && s1[0] != s3[0] {
		return dfs(s1, s2[1:], s3[1:])
	} else if s1[0] == s3[0] || s2[0] == s3[0] {
		return dfs(s1[1:], s2, s3[1:]) || dfs(s1, s2[1:], s3[1:])
	}

	return false
}

//dp
func InterLeavingStrDP(s1, s2, s3 string) bool {

	f := [][]bool{}

	for i := 0; i < len(s2)+1; i++ {
		tmp := make([]bool, len(s1)+1) //行
		f = append(f, tmp)             //列
	}

	f[0][0] = true

	for i := 1; i < len(s1)+1; i++ {
		if !f[i-1][0] || s1[i-1] != s3[i-1] {
			break
		}
		f[i][0] = true
	}

	for i := 1; i < len(s2)+1; i++ {
		if !f[0][i-1] || s2[i-1] != s3[i-1] {
			break
		}
		f[0][i] = true
	}

	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if f[i-1][j] && s1[i-1] == s3[i+j-1] {
				f[i][j] = true
				continue
			}
			if f[i][j-1] && s2[j-1] == s3[i+j-1] {
				f[i][j] = true
			}
		}
	}

	return f[len(s1)][len(s2)]

}
