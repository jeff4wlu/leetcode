package leet121

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBestTimeStock(t *testing.T) {

	Convey("TestBestTimeStock", t, func() {
		Convey("1", func() {
			got := BestTimeStock([]int{6, 2, 9, 8, 1, 3, 7})
			want := 7

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
