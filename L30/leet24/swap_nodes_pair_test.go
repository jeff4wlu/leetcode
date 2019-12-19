package leet24

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSwapNodesPair(t *testing.T) {

	Convey("TestSwapNodesPair", t, func() {
		Convey("常规", func() {

			arr := infra.MakeListNode([]int{1, 2, 3, 4})

			got := SwapNodesPair(arr)
			want := infra.MakeListNode([]int{2, 1, 4, 3})

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
