package leet23

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeKSortedList(t *testing.T) {

	Convey("TestMergeKSortedList", t, func() {
		Convey("常规", func() {
			arr := make([]*infra.ListNode, 3)
			arr[0] = infra.MakeListNode([]int{1, 4, 5})
			arr[1] = infra.MakeListNode([]int{1, 3, 4})
			arr[2] = infra.MakeListNode([]int{2, 6})

			got := MergeKSortedList(arr)
			want := infra.MakeListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
