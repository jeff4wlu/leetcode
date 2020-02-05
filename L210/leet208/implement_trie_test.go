package leet208

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrie(t *testing.T) {

	Convey("TestTrie", t, func() {
		Convey("常规", func() {
			trie := NewTrie()
			trie.Insert("apple")
			got := trie.Search("apple")
			want := true
			if got != want {
				t.Errorf("failed")
			}

			got = trie.Search("app")
			want = false
			if got != want {
				t.Errorf("failed")
			}

			got = trie.StartWith("app")
			want = true
			if got != want {
				t.Errorf("failed")
			}

			trie.Insert("app")
			got = trie.Search("app")
			want = true

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
