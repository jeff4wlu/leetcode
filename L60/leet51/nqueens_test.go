package leet51

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNQueens(t *testing.T) {

	Convey("TestNQueens", t, func() {
		Convey("1", func() {
			got := NQueens(4)
			want := [][]int{
				{1, 3, 0, 2},
				{2, 0, 3, 1},
			}
			if !infra.IntarrCollectionComp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
