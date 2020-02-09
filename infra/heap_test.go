package infra

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHeap(t *testing.T) {

	Convey("TestTopHeap", t, func() {
		Convey("用例1", func() {

			got := NewHeap([]int{3, 9, 20, 15, 7}, true)
			want := []int{3, 7, 20, 15, 9}

			if !IntArrSeqCmp(got.tree, want) {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestBtmHeap", t, func() {
		Convey("用例1", func() {

			got := NewHeap([]int{9, 4, 2, 3, 5, 1, 8, 6, 7, 0}, false)
			want := []int{9, 7, 8, 6, 5, 1, 2, 4, 3, 0}

			if !IntArrSeqCmp(got.tree, want) {
				t.Errorf("failed")
			}

		})

	})

}
