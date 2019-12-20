package leet26

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveDupSortedArr(t *testing.T) {

	Convey("TestRemoveDupSortedArr", t, func() {
		Convey("常规1", func() {

			arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

			got := RemoveDupSortedArr(arr)
			want := []int{0, 1, 2, 3, 4}

			if got != len(want) {
				t.Errorf("got is %d,want is %d", got, len(want))
			} else {
				for i := 0; i < got; i++ {
					if arr[i] != want[i] {
						t.Errorf("got is not equal with want")
					}
				}

			}

		})

		Convey("常规2", func() {

			arr := []int{1, 1}

			got := RemoveDupSortedArr(arr)
			want := []int{1}

			if got != len(want) {
				t.Errorf("got is %d,want is %d", got, len(want))
			} else {
				for i := 0; i < got; i++ {
					if arr[i] != want[i] {
						t.Errorf("got is not equal with want")
					}
				}

			}

		})

		Convey("常规3", func() {

			arr := []int{1, 2}

			got := RemoveDupSortedArr(arr)
			want := []int{1, 2}

			if got != len(want) {
				t.Errorf("got is %d,want is %d", got, len(want))
			} else {
				for i := 0; i < got; i++ {
					if arr[i] != want[i] {
						t.Errorf("got is not equal with want")
					}
				}

			}

		})

	})

}
