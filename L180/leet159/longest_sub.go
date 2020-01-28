package leet159

import "leetcode/infra"

type queue struct {
	c []byte
}

func New() *queue {
	q := new(queue)
	q.c = []byte{}
	return q
}

func (q *queue) enqueue(a byte) {
	q.c = append(q.c, a)
}

func (q *queue) dequeue() byte {
	if len(q.c) > 0 {
		tmp := q.c[0]
		q.c = q.c[1:]
		return tmp
	}
	return ' '
}

func LongestSub(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	max := infra.INT32_MIN
	dict := map[byte]int{}
	q := New()
	for i := 0; i < n; i++ {
		if _, ok := dict[s[i]]; !ok {
			dict[s[i]]++
			q.enqueue(s[i])
		} else {
			dict[s[i]]++
		}

		if len(dict) > 2 {
			var total int
			for _, v := range dict {
				total += v
			}
			max = infra.IntMax(max, total-1)
			tmp := q.dequeue()
			delete(dict, tmp)
		}

	}
	var total int
	for _, v := range dict {
		total += v
	}
	max = infra.IntMax(max, total)

	return max
}
