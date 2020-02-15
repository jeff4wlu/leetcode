package leet264

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUglyNum(t *testing.T) {

	Convey("TestUglyNum", t, func() {
		Convey("用例1", func() {

			got := UglyNum(10)
			want := 12

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

		})

	})

}
