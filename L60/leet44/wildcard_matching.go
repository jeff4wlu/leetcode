package leet44

//s为输入字符串，p是pattern。
//算法主旨是dp[i][j]代表s[0:i-1]与p[0:j-1]是否matching。
//以下算法是从底到顶
func WildcardMatching(s, p string) bool {
	sLen := len(s)
	pLen := len(p)

	//初始化dp,bool默认值是false
	dp := make([][]bool, sLen+1)
	for i := 0; i < sLen+1; i++ {
		dp[i] = make([]bool, pLen+1)
	}

	dp[0][0] = true

	for j := 0; j < pLen; j++ {
		dp[0][j+1] = dp[0][j] && p[j:j+1] == "*"
	}

	for i := 0; i < sLen; i++ {
		for j := 0; j < pLen; j++ {
			if p[j:j+1] != "*" {
				dp[i+1][j+1] = dp[i][j] && (p[j:j+1] == "?" || p[j:j+1] == s[i:i+1])
			} else {
				dp[i+1][j+1] = dp[i][j+1] || dp[i+1][j]
			}
		}
	}

	return dp[sLen][pLen]
}
