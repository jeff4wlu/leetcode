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

	h.BuildHeap()
}

func (h *Heap) BuildHeap() {
	//求最后一个非叶子结点(0base)
	for i := (h.size - 1) / 2; i >= 0; i-- {
		//这里必须使用向下调整，因为向上调整后下面可能又乱了，但调整机会已经过去了。
		h.ShiftDown(i)
	}
}

//向下调整
//确保node是非叶子结点,且只适用于唯一节点值。
//本操作是对单一结点受影响做调整（即默认原来就是堆），而不是全堆
func (h *Heap) ShiftDown(node int) {
	for node <= h.size/2-1 && node >= 0 {
		k := node*2 + 1 //左结点
		if h.isTop {    //小顶堆
			if k+1 < h.size && h.tree[k] > h.tree[k+1] { //求出最小值子结点的位置
				k++
			}
			if h.tree[node] < h.tree[k] {
				break
			}

		} else { //大顶堆
			if k+1 < h.size && h.tree[k] < h.tree[k+1] { //求出最小值子结点的位置
				k++
			}
			if h.tree[node] > h.tree[k] {
				break
			}

		}
		h.tree[node], h.tree[k] = h.tree[k], h.tree[node]
		node = k //由于交换后可能又导致子树不符合堆的条件，所以要调整子树

	}

}

func (h *Heap) ShiftUp(node int) {
	for node <= h.size/2-1 && node >= 0 {
		k := node*2 + 1 //左结点
		if h.isTop {    //小顶堆
			if k+1 < h.size && h.tree[k] > h.tree[k+1] { //求出最小值子结点的位置
				k++
			}
			if h.tree[node] < h.tree[k] {
				break
			}

		} else { //大顶堆
			if k+1 < h.size && h.tree[k] < h.tree[k+1] { //求出最小值子结点的位置
				k++
			}
			if h.tree[node] > h.tree[k] {
				break //不需要调整
			}
		}
		h.tree[node], h.tree[k] = h.tree[k], h.tree[node]
		node = h.parent(node) //由于交换后可能又导致子树不符合堆的条件，所以要调整子树

	}

}

//0base。必须使用它来获取父亲结点，否则返回不了负数
func (h *Heap) parent(idx int) int {

	if idx == 0 {
		return -1
	}

	return (idx - 1) / 2
}

func (h *Heap) GetTop() int {
	if h.size > 0 {
		return h.tree[0]
	}
	return -1
}

//本方法适合处理流数据的优先值
//向上调整
func (h *Heap) Insert(num int) {

	h.tree = append(h.tree, num)
	h.size++

	h.ShiftUp((h.size - 1 - 1) / 2)
}

func (h *Heap) HeapSort(isTop bool) []int {

	res := []int{}
	for h.size > 0 {
		res = append(res, h.tree[0])
		h.tree[0] = h.tree[h.size-1]
		h.size--
		h.BuildHeap()
	}

	return res
}
