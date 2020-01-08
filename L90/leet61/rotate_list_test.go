package leet61

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRotateList(t *testing.T) {

	Convey("TestRotateList", t, func() {
		Convey("1", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 2)
			want := infra.MakeListNode([]int{4, 5, 1, 2, 3})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("超过一个周期", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 7)
			want := infra.MakeListNode([]int{4, 5, 1, 2, 3})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("边界1", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 4)
			want := infra.MakeListNode([]int{2, 3, 4, 5, 1})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})
		Convey("边界2", func() {
			one := infra.MakeListNode([]int{1, 2, 3, 4, 5})
			got := RotateList(one, 5)
			want := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
