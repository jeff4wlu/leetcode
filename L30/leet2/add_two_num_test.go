package leet2

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddTwoNum(t *testing.T) {

	Convey("TestAddTwoNum", t, func() {
		Convey("相同位数相加", func() {
			one := infra.MakeListNode([]int{1, 2, 3})
			two := infra.MakeListNode([]int{1, 2, 3})
			got := AddTwoNum(one, two)
			want := infra.MakeListNode([]int{2, 4, 6})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("相同位数相加，最高位进一", func() {
			one := infra.MakeListNode([]int{9, 9, 9})
			two := infra.MakeListNode([]int{9, 9, 9})
			got := AddTwoNum(one, two)
			want := infra.MakeListNode([]int{8, 9, 9, 1})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("不同位数相加并进位", func() {
			one := infra.MakeListNode([]int{9, 9, 9})
			two := infra.MakeListNode([]int{9, 9})
			got := AddTwoNum(one, two)
			want := infra.MakeListNode([]int{8, 9, 0, 1})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
