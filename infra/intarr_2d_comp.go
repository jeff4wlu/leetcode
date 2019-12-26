package infra

// 比较一个二维阵列是否一样
func IntArr2DComp(got, want [][]int) bool {

	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[0]); j++ {
			if got[i][j] != want[i][j] {
				return false
			}
		}
	}
	return true
}
