package leet41

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstMissingPositive(t *testing.T) {

	Convey("TestFirstMissingPositive", t, func() {
		Convey("1", func() {
			arr := []int{1, 2, 0}
			got := FirstMissingPositive(arr)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {
			arr := []int{3, 4, -1, 1}
			got := FirstMissingPositive(arr)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("3", func() {
			arr := []int{7, 8, 9, 11, 12}
			got := FirstMissingPositive(arr)
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
