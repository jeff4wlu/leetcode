package leet21

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeTwoSortedList(t *testing.T) {

	Convey("TestMergeTwoSortedList", t, func() {
		Convey("相同位数相加", func() {
			one := infra.MakeListNode([]int{1, 2, 4})
			two := infra.MakeListNode([]int{1, 3, 4})
			got := MergeTwoSortedList(one, two)
			want := infra.MakeListNode([]int{1, 1, 2, 3, 4, 4})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

	})

}
