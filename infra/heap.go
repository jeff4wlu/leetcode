package infra

type Heap struct {
	tree  []int
	size  int
	isTop bool
}

//isTop是true时小顶堆
func NewHeap(nums []int, isTop bool) *Heap {
	h := new(Heap)
	h.init(nums, isTop)
	return h
}

func (h *Heap) init(nums []int, isTop bool) {
	h.size = len(nums)
	h.isTop = isTop
	h.tree = make([]int, h.size)
	copy(h.tree, nums)

	if h.isTop {
		h.buildTopHeap()
	} else {
		h.buildBtmHeap()
	}
}

//小顶堆
func (h *Heap) buildTopHeap() {
	//求最后一个非叶子结点(0base)
	for i := h.size/2 - 1; i >= 0; i-- {
		h.adjustTopHeap(i)
	}
}

//大顶堆
func (h *Heap) buildBtmHeap() {
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
