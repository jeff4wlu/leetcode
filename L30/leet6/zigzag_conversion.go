package leet6

func ZigzagConversion(in string, row int) string {

	size := len(in)
	if size < row {
		return ""
	}

	//res := make([]byte, size)
	var res string

	cycleLen := 2*row - 2

	for i := 0; i < row; i++ {
		for j := i; j < size; j += cycleLen {
			res += in[j : j+1]
			pos := j + cycleLen - 2*i
			if i != 0 && i != row-1 && pos < size {
				res += in[pos : pos+1]
			}
		}
	}

	return res

}
