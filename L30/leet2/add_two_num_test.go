package leet2

import (
	"leetcode/base"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNum(t *testing.T) {

	Convey("TestAddTwoNum", t, func() {
		Convey("相同位数相加", func() {
			one := base.MakeListNode([]int{1, 2, 3})
			two := base.MakeListNode([]int{1, 2, 3})
			got := AddTwoNum(one, two)
			want := base.MakeListNode([]int{2, 4, 6})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

		Convey("相同位数相加，最高位进一", func() {
			one := base.MakeListNode([]int{9, 9, 9})
			two := base.MakeListNode([]int{9, 9, 9})
			got := AddTwoNum(one, two)
			want := base.MakeListNode([]int{8, 9, 9, 1})

			for got != nil && want != nil {
				if got.Value != want.Value {
					t.Errorf("got is not equal with want")
					break
				}
				got = got.Next
				want = want.Next
			}

		})

		Convey("不同位数相加并进位", func() {
			one := base.MakeListNode([]int{9, 9, 9})
			two := base.MakeListNode([]int{9, 9})
			got := AddTwoNum(one, two)
			want := base.MakeListNode([]int{8, 9, 0, 1})

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
