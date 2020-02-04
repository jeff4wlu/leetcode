package leet203

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRMLinkedEl(t *testing.T) {

	Convey("TestRMLinkedEl", t, func() {
		Convey("用例1", func() {

			lst := infra.MakeListNode([]int{1, 2, 6, 3, 4, 5, 6})
			got := RMLinkedEl(lst, 6)
			want := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
