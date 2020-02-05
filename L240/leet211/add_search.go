package leet211

type trieNode struct {
	char  byte
	isEnd bool
	nexts []*trieNode
}

//本trie只适用于26个字母
type Trie struct {
	root *trieNode
}

//byte是uint8的别名，所以默认值是0
func newTrieNode(char byte) *trieNode {

	tn := new(trieNode)
	tn.isEnd = false
	tn.char = char
	tn.nexts = make([]*trieNode, 26)
	return tn
}

func NewTrie() *Trie {

	t := new(Trie)
	t.root = newTrieNode('/')
	return t
}

func (t *Trie) AddWord(s string) {

	n := len(s)
	if n == 0 {
		return
	}
	forward := t.root
	for i := 0; i < n; i++ {
		tmp := s[i]
		if forward.nexts[tmp-'a'] == nil {
			forward.nexts[tmp-'a'] = newTrieNode(tmp)
		}
		forward = forward.nexts[tmp-'a']
	}
	forward.isEnd = true
}

//使用dfs收索
func (t *Trie) Search(s string) bool {

	n := len(s)
	if n == 0 {
		return false
	}

	return t.dfs(s, t.root)
}

func (t *Trie) dfs(s string, forward *trieNode) bool {

	n := len(s)
	if n == 0 {
		if forward.isEnd {
			return true
		}
		return false
	}

	for i := 0; i < n; i++ {
		if s[i] == '.' {
			for _, v := range forward.nexts {
				if v != nil && t.dfs(s[i+1:], v) {
					return true
				}
			}
			return false
		}
		if forward.nexts[s[i]-'a'] == nil {
			return false
		}
		forward = forward.nexts[s[i]-'a']
	}
	if forward.isEnd == false {
		return false
	}
	return true
}
