package leet148

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortLst(t *testing.T) {

	Convey("TestSortLst", t, func() {
		Convey("常规", func() {
			one := infra.MakeListNode([]int{4, 2, 1, 3})
			got := SortLst(one)
			want := infra.MakeListNode([]int{1, 2, 3, 4})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("常规1", func() {
			one := infra.MakeListNode([]int{-1, 5, 3, 4, 0})
			got := SortLst(one)
			want := infra.MakeListNode([]int{-1, 0, 3, 4, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
