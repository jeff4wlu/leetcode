package leet277

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindCelebrity(t *testing.T) {

	Convey("TestFindCelebrity", t, func() {

		Convey("1", func() {

			relations := [][]int{
				{1, 1, 0},
				{0, 1, 0},
				{1, 1, 1},
			}

			got := FindCelebrity(relations)
			want := 1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {

			relations := [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{0, 1, 1},
			}

			got := FindCelebrity(relations)
			want := -1

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
