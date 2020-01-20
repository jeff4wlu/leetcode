package leet144

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPreOrder(t *testing.T) {

	Convey("TestPreOrder", t, func() {
		Convey("用例1", func() {

			node1 := new(infra.BTIntNode)
			node2 := new(infra.BTIntNode)
			node3 := new(infra.BTIntNode)
			node1.Value = 1
			node2.Value = 2
			node3.Value = 3
			node1.Right = node2
			node2.Left = node3

			got := PreOrder(node1)
			want := []int{1, 2, 3}

			if !infra.IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
