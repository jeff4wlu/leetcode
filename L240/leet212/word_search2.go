package leet212

type trieNode struct {
	char  byte
	nexts []*trieNode
}

//本trie只适用于26个字母
type Trie struct {
	root *trieNode
}

//byte是uint8的别名，所以默认值是0
func newTrieNode(char byte) *trieNode {

	tn := new(trieNode)
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

	return true
}

func copyBoard(board [][]byte) [][]byte {
	duplicate := make([][]byte, len(board))
	for i := range board {
		duplicate[i] = make([]byte, len(board[i]))
		copy(duplicate[i], board[i])
	}
	return duplicate
}

func WordSearch2(board [][]byte, words []string) []string {

	row := len(board)
	col := len(board[0])
	res := []string{}
	path := []byte{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cpy := copyBoard(board)
			dfs(cpy, i, j, path, &res)
		}
	}

	trie := NewTrie()
	for _, v := range res {
		trie.Insert(v)
	}
	re := []string{}
	for _, v := range words {
		if trie.Search(v) {
			re = append(re, v)
		}
	}
	return re
}

//传入的值要合法的。
func dfs(board [][]byte, x, y int, path []byte, res *[]string) {

	row := len(board)
	col := len(board[0])

	path = append(path, board[x][y])
	board[x][y] = '1'

	//下一步已经无路可走的情况
	if (x-1 < 0 || board[x-1][y] == '1') && (y-1 < 0 || board[x][y-1] == '1') && (x+1 >= row || board[x+1][y] == '1') && (y+1 >= col || board[x][y+1] == '1') {
		(*res) = append((*res), string(path))
		return
	}

	//向上
	if x-1 >= 0 && board[x-1][y] != '1' {
		cpy := copyBoard(board)
		dfs(cpy, x-1, y, path, res)
	}
	//向下
	if x+1 < row && board[x+1][y] != '1' {
		cpy := copyBoard(board)
		dfs(cpy, x+1, y, path, res)
	}
	//向左
	if y-1 >= 0 && board[x][y-1] != '1' {
		cpy := copyBoard(board)
		dfs(cpy, x, y-1, path, res)
	}
	//向右
	if y+1 < col && board[x][y+1] != '1' {
		cpy := copyBoard(board)
		dfs(cpy, x, y+1, path, res)
	}

}
