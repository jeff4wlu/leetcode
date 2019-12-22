package leet34

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstLastSortedArr(t *testing.T) {

	Convey("TestFirstLastSortedArr", t, func() {
		Convey("找到，多个元素", func() {
			arr := []int{5, 7, 7, 8, 8, 10}
			got := FirstLastSortedArr(arr, 8)
			want := []int{3, 4}

			if got[0] != want[0] && got[1] != want[1] {
				t.Errorf("got []int{%d, %d}, want []int{%d, %d}", got[0], got[1], want[0], want[1])
			}

		})
		Convey("找不到", func() {
			arr := []int{5, 7, 7, 8, 8, 10}
			got := FirstLastSortedArr(arr, 6)
			want := []int{-1, -1}

			if got[0] != want[0] && got[1] != want[1] {
				t.Errorf("got []int{%d, %d}, want []int{%d, %d}", got[0], got[1], want[0], want[1])
			}

		})

		Convey("找到，一个元素", func() {
			arr := []int{5}
			got := FirstLastSortedArr(arr, 5)
			want := []int{0, 0}

			if got[0] != want[0] && got[1] != want[1] {
				t.Errorf("got []int{%d, %d}, want []int{%d, %d}", got[0], got[1], want[0], want[1])
			}

		})

	})

}
