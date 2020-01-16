package leet120

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTriangle(t *testing.T) {

	Convey("TestTriangle", t, func() {
		Convey("用例1", func() {

			in := [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			}

			got := Triangle(in)
			want := 11

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
