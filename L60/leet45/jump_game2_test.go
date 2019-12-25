package leet45

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJumpGameII(t *testing.T) {

	Convey("TestJumpGameII", t, func() {
		Convey("1", func() {
			nums := []int{2, 3, 1, 1, 4}

			got := JumpGameII(nums)
			want := 2

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("2", func() {
			nums := []int{2, 3, 1, 1, 4, 2, 3, 1, 3}

			got := JumpGameII(nums)
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
