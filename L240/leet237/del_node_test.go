package leet237

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDelNode(t *testing.T) {

	Convey("TestDelNode", t, func() {
		Convey("常规", func() {
			got := infra.MakeListNode([]int{4, 5, 1, 9})
			DelNode(got.Next)
			want := infra.MakeListNode([]int{4, 1, 9})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

		Convey("常规1", func() {
			got := infra.MakeListNode([]int{4, 5, 1, 9})
			DelNode(got.Next.Next)
			want := infra.MakeListNode([]int{4, 5, 9})

			if !infra.CompareTwoIntList(got, want) {
				t.Errorf("got is not equal with want")
			}

		})

	})

}
