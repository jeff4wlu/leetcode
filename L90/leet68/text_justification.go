package leet68

//没有快捷算法
//扫描单词组，先求出每一行的单词数，再求每一行的空格数。再分配空格。最后处理最后一行的特殊情况
func TextJustification(words []string, maxWidth int) []string {

	wNums := len(words)

	res := []string{}

	var total, startIndex, endIndex int
	for {
		//求每行的words数目
		i := startIndex
		total = 0
		for ; i < wNums; i++ {
			total += len(words[i])
			total++ //空格
			if total-1 > maxWidth {
				break
			}
		}
		endIndex = i
		lineNum := endIndex - startIndex
		if lineNum == 1 { //只有一个单词的行
			var tmp string
			for j := 1; j <= maxWidth-len(words[startIndex]); j++ {
				tmp += " "
			}
			res = append(res, words[startIndex]+tmp)
		} else {
			var spaceNum int
			if endIndex > wNums-1 {
				var tmp string
				for j := startIndex; j < endIndex; j++ {
					tmp += words[j]
					if j < endIndex-1 {
						tmp += " "
					}
				}

				for j := len(tmp) + 1; j <= maxWidth; j++ {
					tmp += " "
				}
				res = append(res, tmp)
				return res

			}
			spaceNum = maxWidth - (total - len(words[endIndex]) - 1) + lineNum

			baseSpace := spaceNum / (lineNum - 1)
			modNum := spaceNum % (lineNum - 1)
			var line string
			if modNum == 0 {
				var tmp string
				for j := 1; j <= baseSpace; j++ {
					tmp += " "
				}
				for j := 0; j < lineNum-1; j++ {
					line += words[startIndex+j]
					line += tmp
				}

			} else {
				var tmp string
				for j := 1; j <= baseSpace; j++ {
					tmp += " "
				}
				k := 1
				for j := 0; j < lineNum-1; j++ {

					line += words[startIndex+j]
					line += tmp
					if k <= modNum {
						line += " "
					}
					k++
				}
			}
			line += words[endIndex-1]
			res = append(res, line)
		}

		if endIndex > wNums-1 {
			break
		}
		startIndex = endIndex

	}

	return res

}
