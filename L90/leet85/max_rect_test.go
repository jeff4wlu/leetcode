package leet85

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxRect(t *testing.T) {

	Convey("TestMaxRect", t, func() {
		Convey("用例1", func() {

			in := [][]int{
				{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0},
			}
			got := MaxRect(in)
			want := 6

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
