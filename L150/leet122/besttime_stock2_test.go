package leet122

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBestTimeStock2(t *testing.T) {

	Convey("TestBestTimeStock2", t, func() {
		Convey("1", func() {
			got := BestTimeStock2([]int{7, 1, 5, 3, 6, 4})
			want := 7

			if got != want {
				t.Errorf("got %d, want  %d", got, want)
			}

		})

	})

}
