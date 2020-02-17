package leet274

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHIndex(t *testing.T) {

	Convey("TestHIndex", t, func() {

		Convey("1", func() {

			got := HIndex([]int{3, 0, 6, 1, 5})
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
