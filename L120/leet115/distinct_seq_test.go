package leet115

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDistinctSeq(t *testing.T) {

	Convey("TestDistinctSeq", t, func() {
		Convey("用例1", func() {

			got := DistinctSeq("rabbbit", "rabbit")
			want := 3

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

		Convey("用例2", func() {

			got := DistinctSeq("babgbag", "bag")
			want := 5

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
