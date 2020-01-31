package leet164

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxGap(t *testing.T) {

	Convey("TestMaxGap", t, func() {
		Convey("用例1", func() {

			got := MaxGap([]int{3, 6, 9, 1})
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := MaxGap([]int{10})
			want := 0

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
