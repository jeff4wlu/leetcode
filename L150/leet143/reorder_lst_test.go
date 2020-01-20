package leet143

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReorderLst(t *testing.T) {

	Convey("TestReorderLst", t, func() {
		Convey("用例1", func() {

			lst := infra.MakeListNode([]int{1, 2, 3, 4})

			got := ReorderLst(lst)
			want := infra.MakeListNode([]int{1, 4, 2, 3})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			lst := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			got := ReorderLst(lst)
			want := infra.MakeListNode([]int{1, 5, 2, 4, 3})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
