package leet279

import (
	"leetcode/infra"
	"math"
)

//dfs是穷举算法，效率极差。
func PerfectSquares(num int) int {

	path := []int{}
	res := [][]int{}
	dfs(num, path, &res)
	min := infra.INT32_MAX
	for _, v := range res {
		min = infra.IntMin(min, len(v))
	}
	if min == infra.INT32_MAX {
		return 0
	}
	return min
}

func dfs(num int, path []int, res *[][]int) {

	if num < 0 {
		return
	}
	if num == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		(*res) = append((*res), tmp)
		return
	}
	n := int(math.Sqrt(float64(num)))

	if n < 1 {
		return
	}
	for i := n; i >= 1; i-- {
		path = append(path, i*i)
		dfs(num-i*i, path, res)
		path = path[:len(path)-1]
	}
}

//求最短路径，即bfs
func PerfectSquares1(num int) int {

	n := int(math.Sqrt(float64(num)))

	count := 1
	q := infra.NewQueue()
	for i := n; i >= 1; i-- {
		diff := num - i*i
		if diff == 0 {
			return count
		}
		if num-i*i > 0 {
			q.Enqueue(i * i)
		}
	}
	if !q.IsEmpty() {
		count++
	}
	q.Enqueue(-1)

	for !q.IsEmpty() {
		a := q.Dequeue().(int)
		if a == -1 {
			if !q.IsEmpty() {
				count++
				q.Enqueue(-1)
				continue
			}
			break
		}
		n := int(math.Sqrt(float64(num - a)))
		for i := n; i >= 1; i-- {
			diff := num - i*i - a
			if diff == 0 {
				return count
			}
			if diff > 0 {
				q.Enqueue(i*i + a)
			}
		}
	}

	return 0
}
