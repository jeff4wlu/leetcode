package leet221

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSquare(t *testing.T) {

	Convey("TestMaxSquare", t, func() {
		Convey("1", func() {

			matrix := [][]int{
				{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 0, 1, 0},
			}

			got := MaxSquare(matrix)
			want := 4

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
