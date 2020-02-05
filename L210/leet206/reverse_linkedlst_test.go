package leet206

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseLinkedLst(t *testing.T) {

	Convey("TestReverseLinkedLst", t, func() {
		Convey("常规", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := ReverseLinkedLst(one)
			want := infra.MakeListNode([]int{5, 4, 3, 2, 1})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("常规1", func() {
			one := infra.MakeListNode([]int{1, 2})
			got := ReverseLinkedLst(one)
			want := infra.MakeListNode([]int{2, 1})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
