package leet173

import (
	"leetcode/infra"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBSTIterator(t *testing.T) {

	Convey("TestBSTIterator", t, func() {
		Convey("1", func() {

			//指针默认值是nil
			tmp := &infra.BTIntNode{}
			tmp.Value = 7
			tmp.Left = &infra.BTIntNode{}
			tmp.Left.Value = 3
			tmp.Right = &infra.BTIntNode{}
			tmp.Right.Value = 15
			tmp.Right.Left = &infra.BTIntNode{}
			tmp.Right.Left.Value = 9
			tmp.Right.Right = &infra.BTIntNode{}
			tmp.Right.Right.Value = 20

			bst := New(tmp)
			got := bst.next()
			want := 3
			if got != want {
				t.Errorf("failed")
			}
			got = bst.next()
			want = 7
			if got != want {
				t.Errorf("failed")
			}

			got1 := bst.hasNext()
			want1 := true
			if got1 != want1 {
				t.Errorf("failed")
			}
			got = bst.next()
			want = 9
			if got != want {
				t.Errorf("failed")
			}
			got1 = bst.hasNext()
			want1 = true
			if got1 != want1 {
				t.Errorf("failed")
			}

			got = bst.next()
			want = 15
			if got != want {
				t.Errorf("failed")
			}
			got1 = bst.hasNext()
			want1 = true
			if got1 != want1 {
				t.Errorf("failed")
			}

			got = bst.next()
			want = 20
			if got != want {
				t.Errorf("failed")
			}
			got1 = bst.hasNext()
			want1 = false
			if got1 != want1 {
				t.Errorf("failed")
			}

		})

	})

}
