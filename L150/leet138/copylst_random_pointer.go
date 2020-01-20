package leet138

type lstNode struct {
	val    int
	next   *lstNode
	random *lstNode
}

func CpyLstRandomPointer(cpy *lstNode) *lstNode {

	dict := make(map[int]*lstNode)

	resNext := new(lstNode)
	res := resNext
	for cpy != nil {

		if _, ok := dict[cpy.val]; !ok {
			resNext.val = cpy.val
			dict[cpy.val] = resNext
		}

		if cpy.next != nil {
			if _, ok := dict[cpy.next.val]; !ok {
				tmp := new(lstNode)
				tmp.val = cpy.next.val
				dict[cpy.next.val] = tmp
			}
			resNext.next = dict[cpy.next.val]
		}

		if cpy.random != nil {
			if _, ok := dict[cpy.random.val]; !ok {
				tmp := new(lstNode)
				tmp.val = cpy.random.val
				dict[cpy.random.val] = tmp
			}
			resNext.random = dict[cpy.random.val]
		}

		cpy = cpy.next
		resNext = resNext.next

	}

	return res
}
