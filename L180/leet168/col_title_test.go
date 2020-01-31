package leet168

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestColTitle(t *testing.T) {

	Convey("TestColTitle", t, func() {
		Convey("用例1", func() {

			got := ColTitle(1)
			want := "A"

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("用例2", func() {

			got := ColTitle(28)
			want := "AB"

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("用例3", func() {

			got := ColTitle(701)
			want := "ZY"

			if got != want {
				t.Errorf("failed.")
			}

		})

	})

}
