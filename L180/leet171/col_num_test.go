package leet171

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColNum(t *testing.T) {

	Convey("TestColNum", t, func() {
		Convey("用例1", func() {

			got := ColNum("C")
			want := 3

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("用例2", func() {

			got := ColNum("AB")
			want := 28

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("用例3", func() {

			got := ColNum("ZY")
			want := 701

			if got != want {
				t.Errorf("failed.")
			}

		})

	})

}
