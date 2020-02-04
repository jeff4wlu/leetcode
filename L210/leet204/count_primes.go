package leet204

import "leetcode/infra"

func CountPrimes(n int) int {

	if n == 0 {
		return 0
	}
	if n == 3 {
		return 1
	}
	if n == 4 {
		return 2
	}

	q := infra.NewQueue()
	q.Enqueue(2)
	q.Enqueue(3)

	res := 2

	dummyHead := new(infra.DListNode)
	pos := dummyHead
	for i := 4; i < n; i++ {
		pos.Next = new(infra.DListNode)
		pos.Next.Value = i
		pos.Next.Pre = pos
		pos = pos.Next
	}

	for q.IsEmpty() == false {
		a := q.Dequeue().(int)
		pos = dummyHead.Next
		for pos != nil {
			b := pos.Value % a
			if b == 0 {
				pos.Pre.Next = pos.Next
				if pos.Next != nil {
					pos.Next.Pre = pos.Pre
				}
			}
			pos = pos.Next
		}
		if dummyHead.Next != nil {
			res++
			q.Enqueue(dummyHead.Next.Value)
			if dummyHead.Next.Next != nil {
				dummyHead.Next.Next.Pre = dummyHead
			}
			dummyHead.Next = dummyHead.Next.Next

		}
	}

	return res
}
