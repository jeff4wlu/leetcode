package leet126

import "leetcode/infra"

import "fmt"

func WordLadder2(start, end string, dict map[string]int) [][]string {

	path := []string{}
	res := [][]string{}

	dfs(start, end, dict, path, &res)

	if len(res) == 0 {
		return [][]string{}
	}

	min := infra.INT32_MAX
	for _, v := range res {
		if len(v) < min {
			min = len(v)
		}
	}

	re := [][]string{}
	for _, v := range res {
		if len(v) == min {
			re = append(re, v)
		}
	}

	return re

}

func dfs(start, end string, dict map[string]int, path []string, res *[][]string) {

	//从字典中找一个与start只相差一个字母的词
	words := []string{}
	for k, v := range dict {
		if v == 1 {
			continue
		}
		l := len(k)
		for i := 0; i < len(k); i++ {
			if start[i] == k[i] {
				l--
			}
		}
		if l == 1 {
			words = append(words, k)
		}
	}

	if len(words) == 0 {
		return
	}

	for _, v := range words {
		if v == end {
			path = append(path, v)
			got := make([]string, len(path))
			copy(got, path)
			*res = append(*res, got)
			return
		}
	}
	fmt.Println()

	for _, v := range words {
		path = append(path, v)
		dict[v] = 1
		dfs(v, end, dict, path, res) //步进
		dict[v] = 0
		path = path[:len(path)-1]
	}

}
