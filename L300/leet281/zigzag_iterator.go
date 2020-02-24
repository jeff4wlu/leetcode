package leet281

import "leetcode/infra"

type zigzagItem struct {
	pos, end, row int
}

type zigzag struct {
	data [][]int
	q    *infra.Queue
}

func NewZigzag(nums [][]int) *zigzag {
	z := new(zigzag)
	z.data = nums
	z.q = infra.NewQueue()
	for i, v := range nums {
		tmp := len(v)
		if tmp > 0 {
			item := new(zigzagItem)
			item.end = tmp
			item.pos = 0
			item.row = i
			z.q.Enqueue(item)
		}
	}
	return z
}

//假设-1是非法值
func (z *zigzag) next() int {
	item := z.q.Dequeue().(*zigzagItem)
	if item.pos+1 < item.end {
		tmp := new(zigzagItem)
		tmp.pos = item.pos + 1
		tmp.end = item.end
		tmp.row = item.row
		z.q.Enqueue(tmp)
	}
	return z.data[item.row][item.pos]
}

func (z *zigzag) hasNext() bool {
	if !z.q.IsEmpty() {
		item := z.q.Peek().(*zigzagItem)
		if item.pos < item.end {
			return true
		}
	}

	return false
}

func ZigzagIterator(nums [][]int) []int {
	z := NewZigzag(nums)
	res := []int{}
	for z.hasNext() {
		res = append(res, z.next())
	}
	return res
}
