package leet25

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverseKGroup(t *testing.T) {

	Convey("TestReverseKGroup", t, func() {
		Convey("常规1", func() {

			arr := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			got := ReverseKGroup(arr, 2)
			want := infra.MakeListNode([]int{2, 1, 4, 3, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("常规2", func() {

			arr := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			got := ReverseKGroup(arr, 3)
			want := infra.MakeListNode([]int{3, 2, 1, 4, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}
		})

	})

}
