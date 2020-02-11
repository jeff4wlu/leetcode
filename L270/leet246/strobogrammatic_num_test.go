package leet246

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStrobogrammaticNum(t *testing.T) {

	Convey("TestStrobogrammaticNum", t, func() {
		Convey("1", func() {

			got := StrobogrammaticNum("69")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("2", func() {

			got := StrobogrammaticNum("88")
			want := true

			if got != want {
				t.Errorf("failed")
			}

		})

		Convey("3", func() {

			got := StrobogrammaticNum("962")
			want := false

			if got != want {
				t.Errorf("failed")
			}

		})

	})

}
