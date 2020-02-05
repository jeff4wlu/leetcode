package leet211

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrie(t *testing.T) {

	Convey("TestTrie", t, func() {
		Convey("常规", func() {
			trie := NewTrie()
			trie.AddWord("bad")
			trie.AddWord("bad")
			trie.AddWord("mad")

			got := trie.Search("pad")
			want := false
			if got != want {
				t.Errorf("failed")
			}

			got = trie.Search("bad")
			want = true
			if got != want {
				t.Errorf("failed")
			}

			got = trie.Search(".ad")
			want = true
			if got != want {
				t.Errorf("failed")
			}
			got = trie.Search("b..")
			want = true
			if got != want {
				t.Errorf("failed")
			}
			got = trie.Search("...")
			want = true
			if got != want {
				t.Errorf("failed")
			}
			got = trie.Search(".az")
			want = false
			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
