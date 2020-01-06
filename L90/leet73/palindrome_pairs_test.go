package leet73

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPalindromePairs(t *testing.T) {

	Convey("TestPalindromePairs", t, func() {
		Convey("用例1", func() {
			strs := []string{"bat", "tab", "cat"}

			got := PalindromePairs(strs)
			want := [][]int{{0, 1}, {1, 0}}

			if got == nil || !infra.IntarrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {
			strs := []string{"abcd", "dcba", "lls", "s", "sssll"}

			got := PalindromePairs(strs)
			want := [][]int{{0, 1}, {1, 0}, {3, 2}, {2, 4}}

			if got == nil || !infra.IntarrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
