package leet42

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrappingRainwater(t *testing.T) {

	Convey("TestTrappingRainwater", t, func() {
		Convey("1", func() {
			arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
			got := TrappingRainwater(arr)
			want := 6

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
