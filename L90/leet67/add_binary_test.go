package leet67

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddBinary(t *testing.T) {

	Convey("TestAddBinary", t, func() {
		Convey("用例1", func() {

			got := AddBinary("11", "1")
			want := "100"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例2", func() {

			got := AddBinary("11", "100")
			want := "111"

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("用例3", func() {

			got := AddBinary("10", "1")
			want := "11"

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
