package leet160

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReorderLst(t *testing.T) {

	Convey("TestReorderLst", t, func() {
		Convey("相交", func() {

			a := infra.MakeListNode([]int{1, 2, 3, 4, 5, 6})

			n1 := new(infra.ListNode)
			n2 := new(infra.ListNode)
			n1.Next = n2
			n2.Next = a.Next.Next.Next
			b := n1

			got := IntersectionLst(a, b)
			want := a.Next.Next.Next

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("不相交", func() {

			a := infra.MakeListNode([]int{1, 2, 3, 4, 5, 6})
			b := infra.MakeListNode([]int{1, 2, 3, 4, 5})

			got := IntersectionLst(a, b)
			var want *infra.ListNode

			if got != want {
				t.Errorf("failed")
			}

		})
	})

}
