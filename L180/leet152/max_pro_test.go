package leet152

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxPro(t *testing.T) {

	Convey("TestMaxPro1", t, func() {
		Convey("1", func() {
			nums := []int{2, 3, -2, 4}
			got := MaxPro1(nums)
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("2", func() {
			nums := []int{-2, 0, -1}
			got := MaxPro1(nums)
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

	Convey("TestMaxPro2", t, func() {
		Convey("1", func() {
			nums := []int{2, 3, -2, 4}
			got := MaxPro2(nums)
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("2", func() {
			nums := []int{-2, 0, -1}
			got := MaxPro2(nums)
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
