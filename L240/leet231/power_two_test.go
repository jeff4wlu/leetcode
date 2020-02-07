package leet231

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPowerTwo(t *testing.T) {

	Convey("TestPowerTwo", t, func() {
		Convey("1", func() {

			got := PowerTwo(1)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := PowerTwo(16)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := PowerTwo(218)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

	Convey("TestPowerTwo1", t, func() {
		Convey("1", func() {

			got := PowerTwo1(1)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := PowerTwo1(16)
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := PowerTwo1(218)
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
