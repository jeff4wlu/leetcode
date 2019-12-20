package leet27

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRemoveElement(t *testing.T) {

	Convey("TestRemoveElement", t, func() {
		Convey("常规1", func() {

			arr := []int{3, 2, 2, 3}

			got := RemoveElement(arr, 3)
			want := []int{2, 2}

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

			arr := []int{0, 1, 2, 2, 3, 0, 4, 2}

			got := RemoveElement(arr, 2)
			want := []int{0, 1, 3, 0, 4}

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
