package leet86

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPartitionList(t *testing.T) {

	Convey("TestPartitionList", t, func() {
		Convey("常规", func() {

			arr := infra.MakeListNode([]int{1, 4, 3, 2, 5, 2})

			got := PartitionList(arr, 3)
			want := infra.MakeListNode([]int{1, 2, 2, 4, 3, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
