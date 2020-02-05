package leet208

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

func (t *Trie) Insert(s string) {

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

func (t *Trie) Search(s string) bool {

	n := len(s)
	if n == 0 {
		return false
	}

	forward := t.root
	for i := 0; i < n; i++ {
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

func (t *Trie) StartWith(s string) bool {

	n := len(s)
	if n == 0 {
		return false
	}
	forward := t.root
	for i := 0; i < n; i++ {
		if forward.nexts[s[i]-'a'] == nil {
			return false
		}
		forward = forward.nexts[s[i]-'a']
	}
	if forward.isEnd == false {
		return true
	}
	return false
}
