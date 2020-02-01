package leet169

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMajorityNum1(t *testing.T) {

	Convey("TestMajorityNum1", t, func() {
		Convey("基数", func() {

			got := MajorityNum1([]int{3, 2, 3})
			want := 3

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("偶数", func() {

			got := MajorityNum1([]int{2, 2, 1, 1, 1, 2, 2, 2})
			want := 2

			if got != want {
				t.Errorf("failed.")
			}

		})

	})

	Convey("TestMajorityNum2", t, func() {
		Convey("基数", func() {

			got := MajorityNum2([]int{3, 2, 3})
			want := 3

			if got != want {
				t.Errorf("failed.")
			}

		})

		Convey("偶数", func() {

			got := MajorityNum2([]int{2, 2, 1, 1, 1, 2, 2, 2})
			want := 2

			if got != want {
				t.Errorf("failed.")
			}

		})

	})

}
