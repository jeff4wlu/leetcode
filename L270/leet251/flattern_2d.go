package leet251

func Flattern2D(nums [][]int) []int {

	v := newVector(nums)
	res := []int{}
	for v.hasNext() {
		res = append(res, v.next())
	}
	return res
}

type vector struct {
	v         [][]int
	x, y, row int //x,y是下一个的坐标，无是-1-1
}

func newVector(nums [][]int) *vector {
	v := new(vector)
	v.v = nums
	v.row = len(nums)
	v.x, v.y = -1, -1

	row := 0
	for ; row < v.row; row++ {
		if len(v.v[row]) > 0 {
			v.x = row
			v.y = 0
			break
		}
	}
	return v
}

func (v *vector) next() int {
	if v.x == -1 {
		return -1
	}
	tmp := v.v[v.x][v.y]

	if v.y+1 < len(v.v[v.x]) {
		v.y++
		return tmp
	}

	row := v.x + 1
	for ; row < v.row; row++ {
		if len(v.v[row]) > 0 {
			v.x = row
			v.y = 0
			return tmp
		}
	}
	v.x, v.y = -1, -1
	return tmp

}

func (v *vector) hasNext() bool {

	if v.x == -1 {
		return false
	}
	return true
}
