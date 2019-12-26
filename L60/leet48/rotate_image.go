package leet48

func RotateImage(img [][]int) {

	n := len(img[0])

	//先上下翻转
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			img[i][j], img[n-1-i][j] = img[n-1-i][j], img[i][j]
		}
	}

	//正对角线翻转
	for i := 0; i < n-1; i++ {
		for j := 1; j < n-i; j++ {
			img[i][i+j], img[i+j][i] = img[i+j][i], img[i][i+j]
		}
	}
}
