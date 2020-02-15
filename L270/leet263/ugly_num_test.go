package leet263

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUglyNum(t *testing.T) {

	Convey("TestUglyNum", t, func() {
		Convey("用例1", func() {

			got := UglyNum(6)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := UglyNum(8)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := UglyNum(14)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
