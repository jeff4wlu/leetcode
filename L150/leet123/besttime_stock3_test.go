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

}
