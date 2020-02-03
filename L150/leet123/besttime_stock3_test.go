package leet123

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBestTimeStock31(t *testing.T) {

	Convey("TestBestTimeStock31", t, func() {
		Convey("1", func() {
			got := BestTimeStock31([]int{3, 3, 5, 0, 0, 3, 1, 4})
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

	Convey("TestBestTimeStock32", t, func() {
		Convey("1", func() {
			got := BestTimeStock32([]int{3, 3, 5, 0, 0, 3, 1, 4})
			want := 6

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("2", func() {
			got := BestTimeStock32([]int{1, 2, 3, 4, 5})
			want := 4

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

		Convey("3", func() {
			got := BestTimeStock32([]int{7, 6, 4, 3, 1})
			want := 0

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
