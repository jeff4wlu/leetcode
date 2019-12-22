package leet35

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchInsertPos(t *testing.T) {

	Convey("TestSearchInsertPos", t, func() {
		Convey("找到", func() {
			arr := []int{1, 3, 5, 6}
			got := SearchInsertPos(arr, 5)
			want := 2

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("找不到", func() {
			arr := []int{1, 3, 5, 6}
			got := SearchInsertPos(arr, 2)
			want := 1

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("找不到，越界右边", func() {
			arr := []int{1, 3, 5, 6}
			got := SearchInsertPos(arr, 7)
			want := 4

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("找不到，越界左边", func() {
			arr := []int{1, 3, 5, 6}
			got := SearchInsertPos(arr, 0)
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
