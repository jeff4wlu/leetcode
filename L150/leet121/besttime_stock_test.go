package leet121

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBestTimeStock(t *testing.T) {

	Convey("TestBestTimeStock", t, func() {
		Convey("1", func() {
			got := BestTimeStock([]int{6, 3, 7, 2, 1, 9, 8})
			want := 8

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
