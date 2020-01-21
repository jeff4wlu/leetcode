package leet146

type LRUCache struct {
	capacity int
	queue    *ItemQueue
	dict     *map[int]*DictNode
}

type DictNode struct {
	value int
	point *ItemNode
}

type LRUCacheAction interface {
	Get(key int) int
	Put(key, value int)
}

func InitLRUCache(capacity int) *LRUCache {
	cache := new(LRUCache)
	cache.capacity = capacity
	cache.queue = New(capacity)
	tmp := map[int]*DictNode{}
	cache.dict = &tmp

	return cache
}

func (s *LRUCache) Get(key int) int {

	if v, ok := (*s.dict)[key]; ok {
		tmp := (*s.dict)[key]
		s.queue.Remove(tmp.point)
		node := s.queue.Enqueue(key)

		tmp.point = node

		return v.value
	}

	return -1
}

// 根据题意简单处理，如果有则离开，不会更新
func (s *LRUCache) Put(key, value int) {
	if _, ok := (*s.dict)[key]; ok {
		return
	}

	node := new(DictNode)
	node.value = value
	//node.point = new(ItemNode)

	(*s.dict)[key] = node
	if s.queue.Size() == s.queue.Cap() {
		tmp := s.queue.Dequeue()
		delete((*s.dict), tmp.Value)
	}
	tmp := s.queue.Enqueue(value)
	node.point = tmp
}
