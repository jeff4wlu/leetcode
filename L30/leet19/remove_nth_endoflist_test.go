package leet19

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveNthEndofList(t *testing.T) {

	Convey("TestRemoveNthEndofList", t, func() {
		Convey("常规", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RemoveNthEndofList(one, 2)
			want := infra.MakeListNode([]int{1, 2, 3, 5})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

		Convey("边界1", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RemoveNthEndofList(one, 5)
			want := infra.MakeListNode([]int{2, 3, 4, 5})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

		Convey("边界2", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RemoveNthEndofList(one, 1)
			want := infra.MakeListNode([]int{1, 2, 3, 4})

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
