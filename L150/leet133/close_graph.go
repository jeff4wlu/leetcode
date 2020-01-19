package leet133

type graphNode struct {
	val       int
	neighbors []*graphNode
}

func CloseGraph(gn *graphNode, dict *map[int]*graphNode) *graphNode {

	newNode := new(graphNode)

	dfs(gn, newNode, dict)
	return newNode
}

func dfs(gn, res *graphNode, dict *map[int]*graphNode) {

	if gn == nil {
		return
	}

	if _, ok := (*dict)[gn.val]; ok {
		return
	}

	res.val = gn.val
	(*dict)[gn.val] = res
	for _, v := range gn.neighbors {
		if _, ok := (*dict)[v.val]; !ok {
			newNode := new(graphNode)
			dfs(v, newNode, dict)
		}
		res.neighbors = append(res.neighbors, (*dict)[v.val])
	}
}
