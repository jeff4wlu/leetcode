package leet82

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDupSortedList(t *testing.T) {

	Convey("TestRemoveDupSortedList", t, func() {
		Convey("常规", func() {

			arr := infra.MakeListNode([]int{1, 2, 3, 3, 4, 4, 5})

			got := RemoveDupSortedList(arr)
			want := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("多于2个的重复值", func() {

			arr := infra.MakeListNode([]int{1, 1, 1, 2, 3})

			got := RemoveDupSortedList(arr)
			want := infra.MakeListNode([]int{1, 2, 3})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
