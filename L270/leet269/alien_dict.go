package leet269

import "leetcode/infra"

//拓扑排序
//dfs找出所有规则；找出所有端点的入度；bfs进行拓扑排序。
func AlienDict(words []string) string {

	if len(words) < 2 {
		return ""
	}

	//求规则
	rules := [][]byte{}
	dfs(words, &rules)
	n := len(rules)
	if n == 0 {
		return ""
	}

	//求入度
	dict := map[byte]int{}
	for _, v := range rules {
		dict[v[1]]++
	}
	for _, v := range rules {
		if _, ok := dict[v[0]]; !ok {
			dict[v[0]] = 0
		}
	}

	q := infra.NewQueue()
	res := []byte{}
	for {
		for k, v := range dict {
			if v < 0 {
				return ""
			}
			if v == 0 {
				q.Enqueue(k)
				delete(dict, k)
			}
		}

		if q.IsEmpty() {
			break
		}

		for !q.IsEmpty() {
			tmp := q.Dequeue().(byte)
			res = append(res, tmp)
			for _, v := range rules {
				if tmp == v[0] {
					dict[v[1]]--
				}
			}
		}
	}

	return string(res)
}

//words是已排序的,递归找出所有rules
func dfs(words []string, rules *[][]byte) {

	n := len(words)
	if n < 2 {
		return
	}

	seq := []byte{}

	var start, i int
	for ; i < n; i++ {
		seq = append(seq, words[i][0])
		if words[i][0] != words[start][0] {
			if i-start > 1 {
				nwords := []string{}
				for j := start; j < i; j++ {
					if len(words[j][1:]) != 0 {
						nwords = append(nwords, words[j][1:])
					}
				}
				if len(nwords) > 1 {
					dfs(nwords, rules)
				}
				start = i
			}
		}

	}
	if i-start > 1 {
		nwords := []string{}
		for j := start; j < i; j++ {
			if len(words[j][1:]) != 0 {
				nwords = append(nwords, words[j][1:])
			}
		}
		if len(nwords) > 1 {
			dfs(nwords, rules)
		}
		start = i
	}

	n = len(seq)
	if n > 1 {
		for i := 1; i < n; i++ {
			if seq[i] != seq[i-1] {
				(*rules) = append((*rules), []byte{seq[i-1], seq[i]})
			}
		}
	}

}
