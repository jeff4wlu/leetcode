package leet170

type twoSum struct {
	data   []int
	length int
}

func New() *twoSum {
	res := new(twoSum)
	res.data = []int{}
	return res
}

func (t *twoSum) add(num int) {
	if t.length == 0 {
		t.data = append(t.data, num)
	} else {
		pos := findPos(t.data, t.length, num)
		t.data = append(t.data[:pos], num)
		t.data = append(t.data, t.data[pos:]...)
	}

	t.length++
}

func (t *twoSum) find(target int) bool {

	if t.length < 2 {
		return false
	}

	for i := 0; i < t.length-1; i++ {
		tmp := target - t.data[i]
		tmp = findP(t.data, t.length, tmp)
		if tmp != -1 {
			return true
		}
	}
	return false
}

//找位置插入
func findPos(in []int, n, target int) int {
	for i := 0; i < n; i++ {
		if target < in[i] {
			return i
		}
	}
	return n
}

//真正查找
func findP(in []int, n, target int) int {
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if in[mid] == target {
			return mid
		}
		if in[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
