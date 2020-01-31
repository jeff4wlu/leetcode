package leet164

import "leetcode/infra"

type bucket struct {
	max   int
	min   int
	empty bool
}

func MaxGap(in []int) int {

	n := len(in)
	if n == 0 {
		return 0
	}
	max := infra.INT32_MIN
	min := infra.INT32_MAX
	for i := 0; i < n; i++ {
		max = infra.IntMax(max, in[i])
		min = infra.IntMin(min, in[i])
	}
	if max == min {
		return 0
	}

	buckets := []*bucket{}
	for i := 0; i <= n; i++ {
		b := new(bucket)
		b.max = infra.INT32_MIN
		b.min = infra.INT32_MAX
		b.empty = true
		buckets = append(buckets, b)
	}

	for i := 0; i < n; i++ {
		idx := put(in[i], n, max, min)
		buckets[idx].empty = false
		buckets[idx].max = infra.IntMax(buckets[idx].max, in[i])
		buckets[idx].min = infra.IntMin(buckets[idx].min, in[i])
	}

	lastMax := buckets[0].max
	var res int
	for i := 1; i <= n; i++ {
		if buckets[i].empty {
			continue
		}
		res = infra.IntMax(res, buckets[i].min-lastMax)
		lastMax = buckets[i].max
	}

	return res
}

func put(v, len, max, min int) int {
	return (v - min) * len / (max - min)
}
