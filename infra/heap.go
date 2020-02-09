package infra

//heap有两种构建方式，一种是对一个已有树构建堆，另外是边读取边构建
type Heap struct {
	tree  []int
	size  int
	isTop bool
}

//isTop是true时小顶堆
func NewHeap1(nums []int, isTop bool) *Heap {
	h := new(Heap)
	h.init(nums, isTop)
	return h
}

func NewHeap2(isTop bool) *Heap {
	h := new(Heap)
	h.size = 0
	h.isTop = isTop
	return h
}

func (h *Heap) init(nums []int, isTop bool) {
	h.size = len(nums)
	h.isTop = isTop
	h.tree = make([]int, h.size)
	copy(h.tree, nums)

	if h.isTop {
		h.BuildTopHeap()
	} else {
		h.BuildBtmHeap()
	}
}

//小顶堆
func (h *Heap) BuildTopHeap() {
	//求最后一个非叶子结点(0base)
	for i := h.size/2 - 1; i >= 0; i-- {
		h.adjustTopHeap(i)
	}
}

//大顶堆
func (h *Heap) BuildBtmHeap() {
	//求最后一个非叶子结点(0base)
	for i := h.size/2 - 1; i >= 0; i-- {
		h.adjustBtmHeap(i)
	}
}

func (h *Heap) adjustTopHeap(node int) {
	if node > h.size/2-1 { //确保被操作的是非叶子结点
		return
	}
	k := node*2 + 1                              //左结点
	if k+1 < h.size && h.tree[k] > h.tree[k+1] { //求出最小值子结点的位置
		k++
	}
	if h.tree[node] > h.tree[k] {
		h.tree[node], h.tree[k] = h.tree[k], h.tree[node]
		h.adjustTopHeap(k) //由于交换后可能又导致子树不符合堆的条件，所以要调整子树
	}

}

func (h *Heap) adjustBtmHeap(node int) {
	if node > h.size/2-1 { //确保被操作的是非叶子结点
		return
	}
	k := node*2 + 1                              //左结点
	if k+1 < h.size && h.tree[k] < h.tree[k+1] { //求出最小值子结点的位置
		k++
	}
	if h.tree[node] < h.tree[k] {
		h.tree[node], h.tree[k] = h.tree[k], h.tree[node]
		h.adjustBtmHeap(k) //由于交换后可能又导致子树不符合堆的条件，所以要调整子树
	}

}

//0base。必须使用它来获取父亲结点，否则返回不了负数
func (h *Heap) parent(idx int) int {

	if idx == 0 {
		return -1
	}

	return (idx - 1) / 2
}

//本方法适合处理流数据的优先值
func (h *Heap) TopInsert(num int) {

	h.tree = append(h.tree, num)
	h.size++

	parent := h.parent(h.size - 1)
	for parent >= 0 {
		k := parent*2 + 1
		if k < h.size {
			if k+1 < h.size && h.tree[k] > h.tree[k+1] {
				k++
			}
			if h.tree[parent] > h.tree[k] {
				h.tree[parent], h.tree[k] = h.tree[k], h.tree[parent]
			}
		}

		parent = h.parent(parent)
	}
}

func (h *Heap) BtmInsert(num int) {

	h.tree = append(h.tree, num)
	h.size++

	parent := h.parent(h.size - 1)
	for parent >= 0 {
		k := parent*2 + 1
		if k < h.size {
			if k+1 < h.size && h.tree[k] < h.tree[k+1] {
				k++
			}
			if h.tree[parent] < h.tree[k] {
				h.tree[parent], h.tree[k] = h.tree[k], h.tree[parent]
			}
		}

		parent = h.parent(parent)
	}
}

func (h *Heap) HeapSort(isTop bool) []int {

	res := []int{}
	for h.size > 0 {
		res = append(res, h.tree[0])
		h.tree[0] = h.tree[h.size-1]
		h.size--
		if isTop {
			h.BuildTopHeap()
		} else {
			h.BuildBtmHeap()
		}
	}

	return res
}
