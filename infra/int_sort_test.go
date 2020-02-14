package infra

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBubbleSort(t *testing.T) {

	Convey("TestBubbleSort", t, func() {
		Convey("排序，从小到大", func() {

			got := BubbleSort([]int{9, 3, 7, 6, 5, 1, 10, 2}, false)
			want := []int{1, 2, 3, 5, 6, 7, 9, 10}

			if !IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

		Convey("排序，从大到小", func() {

			got := BubbleSort([]int{9, 3, 7, 6, 5, 1, 10, 2}, true)
			want := []int{10, 9, 7, 6, 5, 3, 2, 1}

			if !IntArrSeqCmp(got, want) {
				t.Errorf("failed")
			}

		})

	})

}
